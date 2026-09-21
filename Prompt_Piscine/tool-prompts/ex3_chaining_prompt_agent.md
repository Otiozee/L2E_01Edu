# Exercise 3: Chaining Prompts for Agent Flow

## User Query Prompt

### prompt: Ask the user for their query.
"What’s the temperature in Lagos today?"

### Note
This step captures the user’s intent and location. In a real system, the agent would parse the input to determine which tool to use.

## Simulate API Call Prompt
"Simulate a weather API response for the city requested by the user.
Return the output in valid JSON format using this schema:

{
  "location": "string",
  "temperature_celsius": number,
  "condition": "string"
}

Input:
City = Lagos"

### model response
```json
{
  "location": "Lagos",
  "temperature_celsius": 31,
  "condition": "Sunny"
}
```
### Note
At this stage, the agent produces structured data just like an actual API would. It’s machine readable and precise.

## Format Final Answer Prompt
"Take the JSON from the simulated API and format a human readable response for the user.

Input:

{
  "location": "Lagos",
  "temperature_celsius": 31,
  "condition": "Sunny"
}"

### Model output
```markdown
The current temperature in Lagos is 31°C and sunny.
```

## Chained Process review
- User Query Step: Captures the user’s request and extracts relevant information (city name).

- Tool Simulation Step: Generates structured output (JSON) that mimics a real weather API.

- Formatting Step: Converts structured data into a human-readable answer for the user.

## Reflection
The chain ensures consistency since the data flows from input to tool to output without losing accuracy. Using JSON in the middle step allows programmatic access if needed. The final answer feels natural to the user, while still being based on structured data.