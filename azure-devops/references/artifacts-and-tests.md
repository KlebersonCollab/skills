# Artifacts & Test Plans: Azure DevOps

This guide covers package management (Azure Artifacts) and test orchestration (Azure Test Plans).

## 1. Azure Artifacts (Packages)

Artifacts allow creating private feeds to share packages among teams. Supports NuGet, npm, Maven, Python, and Universal Packages.

### List Feeds
`GET https://feeds.dev.azure.com/{org}/{project}/_apis/packaging/feeds?api-version=7.1`

### List Package Versions
`GET https://feeds.dev.azure.com/{org}/{project}/_apis/packaging/Feeds/{feedId}/Packages/{packageId}/versions?api-version=7.1`

## 2. Azure Test Plans (Tests)

Manages the entire life cycle of manual and automated tests.

### Plans and Suites
- **Test Plan**: Top-level container for tests.
- **Test Suite**: Logical grouping of test cases within a plan.
- **Test Case**: Individual test steps.

### List Test Plans
`GET https://dev.azure.com/{org}/{project}/_apis/test/plans?api-version=7.1`

## 3. Test Execution

Record results of a test execution:

`POST https://dev.azure.com/{org}/{project}/_apis/test/runs?api-version=7.1`

Payload:
```json
{
  "name": "Automated Test Execution",
  "plan": { "id": "42" },
  "isAutomated": true
}
```

## 4. Pipeline Integration

Azure Test Plans can be integrated into pipelines to publish results from unit or integration tests:

```yaml
- task: PublishTestResults@2
  inputs:
    testResultsFormat: 'JUnit'
    testResultsFiles: '**/TEST-*.xml'
```

## 5. Best Practices

- **Immutable Versioning**: Never replace an already published package version. Always upload a new version.
- **Upstream Sources**: Use upstream sources in Azure Artifacts to centralize the consumption of public packages (e.g., npmjs.org) and private packages in a single feed.
- **Test Case Reusability**: Use Shared Steps in Test Cases to avoid instruction duplication.
- **Traceability**: Always link test results to Work Items and Pull Requests to ensure full traceability.
