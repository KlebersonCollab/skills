# MCP Interaction Modes

Production agents need more than just text. MCP provides several primitives for richer user interaction.

## 1. Inline UI (MCP Apps)
Allows the server to return an interactive interface that the client renders inline in chat.
- **Visuals**: Charts, tables, maps.
- **Interactions**: Buttons, toggles, mini-dashboards.
- **Constraint**: Requires the client to support the specific component framework or follow the MCP Apps spec.

## 2. Elicited Input (Form Mode)
The server pauses a tool execution and requests a structured form from the user.
- **Flow**:
  1. Agent calls tool.
  2. Server responds with an elicitation request (Form schema).
  3. Client renders the form.
  4. User fills and submits.
  5. Control returns to the server with the data.
- **Best Practices**: Use for ambiguity resolution and high-stakes confirmations.

## 3. External Handoff (URL Mode)
The server provides a URL that the client opens in the user's browser.
- **Security**: The "Sensitive Data" never reaches the LLM or the Chat interface.
- **Flow**:
  1. Server returns a URL Mode elicitation.
  2. Browser opens.
  3. User completes task (e.g., Payment).
  4. Server receives webhook/callback.
  5. Flow resumes in chat.

## 4. Prompt Elicitation
A simpler form of interaction where the server provides a pre-filled prompt for the user to confirm or edit before it hits the model.
