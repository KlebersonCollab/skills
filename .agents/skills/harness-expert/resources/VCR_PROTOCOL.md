# Session VCR Protocol (Recording & Replay)

This protocol defines how to record and replay agent sessions for debugging and behavior validation.

## 1. Recording a Session
When the user or a mandate requests a "VCR Recording":
1.  **Initialize Trace**: Create a JSON file in `.specs/history/vcr/<task-id>-<timestamp>.json`.
2.  **Capture Turn**: For every turn, record:
    - `user_prompt`: The original input.
    - `thought`: The agent's thinking block.
    - `tool_calls`: All tool invocations and their arguments.
    - `tool_outputs`: The raw output from the system.
    - `assistant_response`: The final message sent to the user.

## 2. The Golden Path (Baseline)
A "Golden Path" is a VCR recording of a task performed perfectly. 
- **Requirement**: Must be verified by a human or the `sdd-reviewer`.
- **Use Case**: Used as a reference to detect regression in new versions of skills.

## 3. Replay & Analysis (Rewind)
To analyze a failure using a VCR trace:
1.  **Load Trace**: Read the JSON file.
2.  **Turn-by-Turn Comparison**:
    - Turn X (VCR): Agent calls `read_file(A)`.
    - Turn X (Live): Agent calls `ls(B)`.
    - **Divergence Point**: This is where the logic failed. Investigate the `thought` block in the live session vs. the VCR.

## 4. Portability (Hydration)
- **CWD Normalization**: Replace absolute paths with `[CWD]` in the trace to allow replay on different machines.
- **Timestamp Normalization**: Replace real dates with `[TIMESTAMP]` to avoid logic branching based on time.
