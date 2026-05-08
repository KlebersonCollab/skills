const fs = require('fs');
const path = require('path');

const SKILLS_DIR = path.join(__dirname, '../../.agents/skills');
const GLOBAL_MANDATES_PATH = path.join(__dirname, '../../.specs/codebase/GLOBAL_MANDATES.md');
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

  // First pass: gather all metadata and content
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

  // Second pass: detect relationships (how they "talk" to each other)
  const skillIds = skills.map(s => s.id);
  
  // Parse Global Mandates for mandatory routing
  let globalMandateConns = [];
  if (fs.existsSync(GLOBAL_MANDATES_PATH)) {
    const mandates = fs.readFileSync(GLOBAL_MANDATES_PATH, 'utf8');
    // Extract skill IDs from the Skill Router table
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
    
    // If it's a global mandate, it talks to SDD by default (orchestration)
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

// Ensure scripts and public dir exists
const scriptsDir = path.join(__dirname);
if (!fs.existsSync(scriptsDir)) {
  fs.mkdirSync(scriptsDir, { recursive: true });
}
const publicDir = path.join(__dirname, '../public');
if (!fs.existsSync(publicDir)) {
  fs.mkdirSync(publicDir, { recursive: true });
}

const data = scanSkills();
fs.writeFileSync(OUTPUT_FILE, JSON.stringify({ skills: data }, null, 2));
console.log(`Registry updated: ${data.length} skills found.`);
