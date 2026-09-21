# Exercise 2: Simulated API Call (Weather Example)

## Basic Prompt
"Simulate a weather API response for Paris today."

### Model response
```markdown
The weather in Paris today is 18°C and Cloudy.
```
### Personal note:
This is readable for a human, but it is not structured. If I want to use it in a program or system, I would need to extract the data manually.

## Improved Prompt With Defined JSON Schema
"Simulate a weather API response for Paris today.
Return the result in valid JSON format using the following schema ("location": "string","temperature_celsius": number,"condition": "string"). Do not include any extra text outside the JSON."

### Model response
```json
{
  "location": "Paris",
  "temperature_celsius": 18,
  "condition": "Cloudy"
}
```
### Personal note
The output now looks like a real API response. It is structured, machine readable and easy to use in software.

## Testing With Other Cities: Abuja
"Simulate a weather API response for Abuja today.
Return the result in valid JSON format using the same schema. Do not include extra text."

### Model response
```json
{
  "location": "Abuja",
  "temperature_celsius": 28,
  "condition": "Partly Cloudy"
}
```
## Testing With Other Cities: Port-Harcourt
"Simulate a weather API response for Port Harcourt today.
Return the result in valid JSON format using the same schema. Do not include extra text."

### model response
```json
{
  "location": "Port Harcourt",
  "temperature_celsius": 33,
  "condition": "Rainy"
}
```
## Comparing output and final reflection
- When the schema is defined, the model outputs structured JSON consistently, and only the values change for each city. Without the schema, the model defaults to general sentences, which are harder to use in programming.
- Using a clear JSON schema in the prompt is crucial when you want the model to behave like an API. This ensures consistency, machine readability, and makes it easier to integrate into applications, whether for Lagos, Abuja or any other city.
