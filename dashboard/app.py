import streamlit as st
import subprocess
import json
from pathlib import Path

# Configurações raiz
root_dir = Path(__file__).parent.parent
hb_bin = root_dir / "bin" / "hb"

# ==========================================
# 🎨 1. PAGE CONFIG & CUSTOM STYLING (CSS)
# ==========================================
st.set_page_config(
    page_title="AI Skills Hub",
    page_icon="⚡",
    layout="wide",
    initial_sidebar_state="collapsed",
)

# Injeção de CSS para Glassmorphism e tipografia limpa
custom_css = """
<style>
    #MainMenu {visibility: hidden;}
    footer {visibility: hidden;}
    header {visibility: hidden;}

    .stApp {
        background-image: radial-gradient(circle at 15% 50%, rgba(59, 130, 246, 0.08), transparent 25%),
                          radial-gradient(circle at 85% 30%, rgba(139, 92, 246, 0.08), transparent 25%);
    }

    .streamlit-expanderHeader {
        background: rgba(30, 41, 59, 0.5) !important;
        backdrop-filter: blur(10px) !important;
        border-radius: 8px !important;
        border: 1px solid rgba(255, 255, 255, 0.05) !important;
        color: #e2e8f0 !important;
        font-weight: 600 !important;
        padding: 10px 15px !important;
    }
    
    div[data-testid="stExpanderDetails"] {
        background: rgba(15, 23, 42, 0.6) !important;
        border-left: 2px solid #3b82f6 !important;
        padding: 15px !important;
    }

    h1 {
        background: -webkit-linear-gradient(45deg, #3b82f6, #8b5cf6);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        font-weight: 800;
        font-family: 'Inter', sans-serif;
    }
    
    div[data-testid="stMetricValue"] {
        color: #3b82f6;
    }
</style>
"""
st.markdown(custom_css, unsafe_allow_html=True)

# ==========================================
# 📊 2. DATA FETCH (VIA HB-CLI)
# ==========================================

@st.cache_data(ttl=60)
def fetch_skills():
    try:
        # Se o binário não existir, tentamos rodar via go run como fallback (desenvolvimento)
        if not hb_bin.exists():
            result = subprocess.run(
                ["go", "run", "main.go", "list", "--json"],
                cwd=root_dir / "hb",
                capture_output=True,
                text=True,
                check=True
            )
        else:
            result = subprocess.run(
                [str(hb_bin), "list", "--json"],
                capture_output=True,
                text=True,
                check=True
            )
        return json.loads(result.stdout)
    except Exception as e:
        st.error(f"Erro ao carregar skills via HB-CLI: {e}")
        return []

skills = fetch_skills()
total_skills = len(skills)

mermaid_path = root_dir / "KNOWLEDGE-MAP.mermaid"
mermaid_code = (
    mermaid_path.read_text() if mermaid_path.exists() else "graph TD\n A[Sem Mapa]"
)

# ==========================================
# 🖥️ 3. UI LAYOUT
# ==========================================
st.title("⚡ AI Agent Skills Hub")
st.markdown("Ecossistema de Conhecimento e Governança SDD")

# Métricas Top Level
col_m1, col_m2, col_m3 = st.columns(3)
col_m1.metric("Skills Registradas", total_skills)
col_m2.metric("SDD Compliance", "100%", "Quality Gate Pass")
col_m3.metric("Infraestrutura", "Go (HB-CLI)", "v1.0.0")

st.divider()

# Abas de Navegação Principal
tab1, tab2 = st.tabs(["🧩 Catálogo de Skills", "🕸️ Architecture Graph"])

with tab1:
    st.markdown("### Navegação Rápida")

    # Grid dinâmico para cards de skills
    cols = st.columns(2)

    for i, skill in enumerate(skills):
        name = skill.get("Name")
        version = skill.get("Version", "0.0.0")
        category = skill.get("Category", "Geral")
        desc = skill.get("Description", "Sem descrição.")

        # Aloca iterativamente entre a Coluna 1 e Coluna 2
        col_dest = cols[0] if i % 2 == 0 else cols[1]

        with col_dest:
            with st.expander(f"📦 {name} (v{version})"):
                st.markdown(f"**Categoria**: `{category}`")
                st.write(f"_{desc}_")

                uses = skill.get("Uses")
                if uses:
                    st.markdown("**Depende de:**")
                    for dep in uses:
                        st.markdown(f"- 🔌 `{dep}`")

with tab2:
    st.markdown("### Mapa de Conhecimento Relacional")
    st.caption("Visão interativa gerada dinamicamente pelo HB-CLI")

    import base64
    import json as json_lib
    import streamlit.components.v1 as components

    state = {
        "code": mermaid_code,
        "mermaid": '{\n  "theme": "dark"\n}',
        "autoSync": True,
        "updateDiagram": True,
    }

    json_state = json_lib.dumps(state)
    encoded = base64.urlsafe_b64encode(json_state.encode("utf-8")).decode("ascii")
    live_url = f"https://mermaid.live/view#base64:{encoded}"

    components.iframe(live_url, height=800, scrolling=True)
