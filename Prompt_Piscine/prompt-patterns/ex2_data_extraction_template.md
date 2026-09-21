# Exercise 2: Data Extraction Template

## Template Prompt:
Extract the following fields from the text: Name, Age, Location, Occupation, employer, any other relevant dates(birth_date, Joined_date). Return the output in JSON format.
Rules:
- If a field is missing or unclear return null  
- Return ONLY valid JSON, no extra explanation  
- Use snake_case keys  
- Do not invent or assume information

### Test Input 1: John Doe, age 29, lives in Paris and works as a software engineer.
Model output:
```json
{
  "name": "John Doe",
  "age": 29,
  "location": "Paris",
  "occupation": "software engineer",
  "employer": null
}
```
### Test input 2: Maria Silva (born 15 March 1992) is a senior product manager at Stripe in Lisbon, Portugal. She joined the company in 2020.
Model output:
```json
{
  "name": "Maria Silva",
  "age": null,
  "location": "Lisbon, Portugal",
  "occupation": "senior product manager",
  "employer": "Stripe",
  "birth_date": "15 March 1992",
  "joined_date": "2020"
}
```
### Test input 3: The new head of marketing is Alex Chen.
Model output:
```json
{
  "name": "Alex Chen",
  "age": null,
  "location": null,
  "occupation": "head of marketing",
  "employer": null
}
```
## Analysis
The template enforces strict JSON output, prevents hallucination of missing fields, consistently uses null instead of empty strings or guesses, and allows graceful extension for new fields. Very reliable across short and longer inputs.