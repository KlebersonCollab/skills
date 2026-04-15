# Artifacts & Test Plans: Azure DevOps

Este guia cobre a gestão de pacotes (Azure Artifacts) e a orquestração de testes (Azure Test Plans).

## 1. Azure Artifacts (Pacotes)

Artifacts permitem criar feeds privados para compartilhar pacotes entre times. Suporta NuGet, npm, Maven, Python e Universal Packages.

### Listar Feeds
`GET https://feeds.dev.azure.com/{org}/{project}/_apis/packaging/feeds?api-version=7.1`

### Listar Versões de um Pacote
`GET https://feeds.dev.azure.com/{org}/{project}/_apis/packaging/Feeds/{feedId}/Packages/{packageId}/versions?api-version=7.1`

## 2. Azure Test Plans (Testes)

Gerencia todo o ciclo de vida de testes manuais e automatizados.

### Planos e Suítes
- **Test Plan**: Container de nível superior para testes.
- **Test Suite**: Agrupamento lógico de casos de teste dentro de um plano.
- **Test Case**: Passos individuais de teste.

### Listar Planos de Teste
`GET https://dev.azure.com/{org}/{project}/_apis/test/plans?api-version=7.1`

## 3. Execução de Testes

Registrar resultados de uma execução de teste:

`POST https://dev.azure.com/{org}/{project}/_apis/test/runs?api-version=7.1`

Payload:
```json
{
  "name": "Execução de Teste Automatizado",
  "plan": { "id": "42" },
  "isAutomated": true
}
```

## 4. Integração com Pipelines

Azure Test Plans pode ser integrado aos pipelines para publicar resultados de testes unitários ou de integração:

```yaml
- task: PublishTestResults@2
  inputs:
    testResultsFormat: 'JUnit'
    testResultsFiles: '**/TEST-*.xml'
```

## 5. Boas Práticas

- **Versionamento Imutável**: Nunca substitua uma versão de pacote já publicada. Sempre suba uma nova versão.
- **Upstream Sources**: Utilize upstream sources no Azure Artifacts para centralizar o consumo de pacotes públicos (ex: npmjs.org) e pacotes privados em um único feed.
- **Test Case Reusability**: Utilize Shared Steps em Test Cases para evitar duplicação de instruções.
- **Traceability**: Sempre vincule resultados de teste a Work Items e Pull Requests para garantir rastreabilidade completa.
