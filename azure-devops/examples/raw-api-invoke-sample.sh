#!/bin/bash

# Exemplo 1: Listar logs de auditoria da organização (API audit)
# Útil quando o comando padrão da CLI não atende à granularidade necessária.
echo "Invocando API de Auditoria..."
az devops invoke \
    --area "audit" \
    --resource "auditlog" \
    --route-values project="MeuProjeto" \
    --http-method GET \
    --output json

# Exemplo 2: Listar membros de um grupo de segurança (API identity)
# Note que o 'subjectDescriptor' deve ser obtido previamente.
echo "Consultando identidades do grupo..."
az devops invoke \
    --area "graph" \
    --resource "memberships" \
    --route-values subjectDescriptor="vssgp.Uy0xLTktMTU1MTM3NDI0NS0yMzQwMTM5..." \
    --http-method GET \
    --output json

# Exemplo 3: Listar métricas de um Dashboard
echo "Consultando widgets do dashboard..."
az devops invoke \
    --area "dashboard" \
    --resource "widgets" \
    --route-values project="MeuProjeto" dashboardId="uuid-do-dashboard" \
    --http-method GET \
    --output json
