# Gold Standard Scenario: Knowledge Mapping

**Problem**: The payment system is getting complex. The agent needs to map how the new "Refund" requirements connect to legacy classes and the API.

**Input**: "Map the relationships of the Refund feature with the current code using the knowledge-architect."

**Execution**:
1.  **Exploration**: The agent reads the existing code for payments and the new spec.
2.  **Mapping**: Identifies that `RefundService` depends on `LegacyPaymentGateway` and updates `TransactionTable`.
3.  **Visualization**: Generates a Mermaid graph showing these connections.
4.  **Verification**: Ensures that all critical dependencies are represented.
