"""
{{project_name}}
{{'=' * project_name|length}}

{{description}}
"""

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

{% if database == 'postgresql' %}
from {{module_name}}.database import engine
from {{module_name}}.models import Base
{% endif %}

app = FastAPI(
    title="{{project_name}}",
    description="{{description}}",
    version="0.1.0",
)

# CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

{% if database == 'postgresql' %}
# Create tables
Base.metadata.create_all(bind=engine)
{% endif %}

{% if include_auth %}
from {{module_name}}.auth import router as auth_router
app.include_router(auth_router, prefix="/api/auth", tags=["auth"])
{% endif %}

@app.get("/")
async def root():
    """Root endpoint"""
    return {
        "message": "Welcome to {{project_name}}",
        "version": "0.1.0",
    }

@app.get("/health")
async def health():
    """Health check endpoint"""
    return {"status": "healthy"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)