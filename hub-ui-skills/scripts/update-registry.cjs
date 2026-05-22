const fs = require('fs');
const path = require('path');

const SKILLS_DIR = path.join(__dirname, '../../.agents/skills');
const GLOBAL_MANDATES_PATH = path.join(__dirname, '../../.specs/codebase/GLOBAL_MANDATES.md');
const HARNESS_DIR = path.join(__dirname, '../../.harness');
const OUTPUT_FILE = path.join(__dirname, '../public/registry.json');

function parseMarkdown(content) {
  const meta = {};
  const frontmatterMatch = content.match(/^---\n([\s\S]*?)\n---/);
  if (frontmatterMatch) {
    frontmatterMatch[1].split('\n').forEach(line => {
      const [key, ...value] = line.split(':');
      if (key && value) {
        meta[key.trim()] = value.join(':').trim().replace(/^["']|["']$/g, '');
      }
    });
  }

  const sddMatch = content.match(/<!-- @sdd-state -->\n```yaml\n([\s\S]*?)\n```/);
  if (sddMatch) {
    const sdd = {};
    sddMatch[1].split('\n').forEach(line => {
      const [key, ...value] = line.split(':');
      if (key && value) {
        sdd[key.trim()] = value.join(':').trim().replace(/^["']|["']$/g, '');
      }
    });
    meta.sddState = sdd;
  }

  return meta;
}

function scanSkills() {
  const skills = [];
  const dirs = fs.readdirSync(SKILLS_DIR).filter(d => fs.statSync(path.join(SKILLS_DIR, d)).isDirectory());

  dirs.forEach(dir => {
    const skillPath = path.join(SKILLS_DIR, dir, 'SKILL.md');
    if (fs.existsSync(skillPath)) {
      const content = fs.readFileSync(skillPath, 'utf8');
      const meta = parseMarkdown(content);
      skills.push({
        id: dir,
        ...meta,
        content: content,
        path: `.agents/skills/${dir}/SKILL.md`
      });
    }
  });

  const skillIds = skills.map(s => s.id);
  
  let globalMandateConns = [];
  if (fs.existsSync(GLOBAL_MANDATES_PATH)) {
    const mandates = fs.readFileSync(GLOBAL_MANDATES_PATH, 'utf8');
    const tableMatch = mandates.match(/## 📍 SKILL ROUTER[\s\S]*?\|([\s\S]*?)\n\n/);
    if (tableMatch) {
      const rows = tableMatch[1].split('\n').filter(r => r.includes('`'));
      rows.forEach(row => {
        const skillsInRow = row.match(/`([^`]+)`/g);
        if (skillsInRow) {
          skillsInRow.forEach(s => {
            const id = s.replace(/`/g, '').split(' ')[0].trim();
            if (skillIds.includes(id)) globalMandateConns.push(id);
          });
        }
      });
    }
  }

  skills.forEach(skill => {
    skill.conversesWith = [];
    
    if (globalMandateConns.includes(skill.id) && skill.id !== 'sdd') {
      skill.conversesWith.push('sdd');
    }

    skillIds.forEach(otherId => {
      if (otherId === skill.id) return;
      
      const regex = new RegExp(`\\b${otherId}\\b`, 'i');
      if (regex.test(skill.content) && !skill.conversesWith.includes(otherId)) {
        skill.conversesWith.push(otherId);
      }
    });
  });

  return skills;
}

function scanHarness() {
  const harness = {
    active_provider: 'unknown',
    providers: {},
    sessions: [],
    session_count: 0,
    last_session: null
  };

  // Read config
  const configPath = path.join(HARNESS_DIR, 'config.json');
  if (fs.existsSync(configPath)) {
    try {
      const config = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      harness.active_provider = config.active_provider || 'unknown';
      harness.providers = config.providers || {};
    } catch (e) {
      console.error('Error reading harness config:', e.message);
    }
  }

  // Read sessions
  const sessionsDir = path.join(HARNESS_DIR, 'sessions');
  if (fs.existsSync(sessionsDir)) {
    try {
      const sessionFiles = fs.readdirSync(sessionsDir).filter(f => f.endsWith('.json'));
      harness.session_count = sessionFiles.length;
      
      sessionFiles.forEach(f => {
        const sessionPath = path.join(sessionsDir, f);
        try {
          const session = JSON.parse(fs.readFileSync(sessionPath, 'utf8'));
          // Extract summary
          const summary = {
            session_id: session.session_id,
            root_task: session.root_task?.substring(0, 80) || '',
            active_node: session.active_node || '',
            node_count: Object.keys(session.nodes || {}).length,
            timestamp: session.nodes?.[Object.keys(session.nodes || {})[0]]?.timestamp || ''
          };
          harness.sessions.push(summary);
        } catch (e) {
          console.error(`Error parsing session ${f}:`, e.message);
        }
      });

      // Sort sessions by timestamp descending
      harness.sessions.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp));
      
      if (harness.sessions.length > 0) {
        harness.last_session = harness.sessions[0];
      }
    } catch (e) {
      console.error('Error reading sessions dir:', e.message);
    }
  }

  return harness;
}

// Ensure directories exist
const scriptsDir = path.join(__dirname);
if (!fs.existsSync(scriptsDir)) {
  fs.mkdirSync(scriptsDir, { recursive: true });
}
const publicDir = path.join(__dirname, '../public');
if (!fs.existsSync(publicDir)) {
  fs.mkdirSync(publicDir, { recursive: true });
}

// Copy session files to public/sessions/ for client-side fetch
const sessionsDir = path.join(HARNESS_DIR, 'sessions');
const publicSessionsDir = path.join(publicDir, 'sessions');
if (!fs.existsSync(publicSessionsDir)) {
  fs.mkdirSync(publicSessionsDir, { recursive: true });
}
if (fs.existsSync(sessionsDir)) {
  fs.readdirSync(sessionsDir).filter(f => f.endsWith('.json')).forEach(f => {
    fs.copyFileSync(path.join(sessionsDir, f), path.join(publicSessionsDir, f));
  });
}

const skills = scanSkills();
const harness = scanHarness();

const data = { skills, harness };
fs.writeFileSync(OUTPUT_FILE, JSON.stringify(data, null, 2));
console.log(`Registry updated: ${data.skills.length} skills, ${data.harness.session_count} harness sessions found.`);
