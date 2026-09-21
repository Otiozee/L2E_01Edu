# Exercise 1: Structured JSON Output

## Unstructured Prompt
"List three cities with their populations."

### Model response
```markdown
Tokyo – about 37 million people
Delhi – around 32 million
Shanghai – roughly 29 million
```
### Personal note
This answer is readable, but it is not structured in a consistent way. The words (“about,” “around,” “roughly”) makes it harder to use in a program. If I wanted to automatically process this data, I would need extra steps to clean and standardize it.

## Improved Prompt With JSON Structure
"List three cities with their populations.
Return the result in valid JSON format.
Do not include any text outside the JSON."

### Model response
```json
[
  { "city": "Tokyo", "population": 37000000 },
  { "city": "Delhi", "population": 32000000 },
  { "city": "Shanghai", "population": 29000000 }
]
```
### Personal note
Now the output is well structured and machine readable. The population values are numbers instead of text with approximations. This format can easily be parsed by a program.

## Testing With a Different Input
"List three cities in Nigeria with their approximate populations.
Return the result in valid JSON format.
Do not include any text outside the JSON."

### Model response
```json
[
  { "city": "Lagos", "population": 15000000 },
  { "city": "Kano", "population": 4000000 },
  { "city": "Ibadan", "population": 3500000 }
]
```
## Reflection
When prompts are vague, the results are written for humans and may vary in format. By clearly specifying a JSON structure and restricting extra text, the output becomes predictable and reliable. This is especially useful when building systems that depend on structured data rather than plain text. Overall, the output follows the exact JSON structure requested and each object contains a string for the city name. Also, each population value is a number, not text and there is no extra explanation outside the JSON block.