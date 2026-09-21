# Exercise 2: Adding Source or Data Constraints

## Factual Question
"What year was the company founded?"

### Analysis
This question is clear, but if no source is provided, the model might guess or pull information from somewhere else. That increases the risk of incorrect answers.

## Strengthened Prompt with Constraints
"Read the passage below and answer the question.
Use only the information provided in the text.
Cite the exact sentence from the passage in your answer.

Passage:
RunGoTech Solutions was founded in 2024 by two engineering graduates. The company started as a small startup focused on solar panel installation. By 2025, it had expanded into energy storage systems and smart grid technology. Today, RunGoTech operates in more than 15 countries.

Question:
When was RunGoTech Solutions founded?"

### Model response
```markdown
RunGoTech Solutions was founded in 2024.

Source: “RunGoTech Solutions was founded in 2024 by two engineering graduates.”
```
## Reflection
- In the first version, the question is too open. Without a source, the model could rely on outside knowledge or even guess and or fabricate.

- In the improved version, the instructions clearly limit the answer to the given passage. Asking for a citation forces the response to stay grounded in the text. The answer is short, accurate, and directly supported by the provided material.

- overall, this approach reduces hallucination and makes it easier to verify the correctness of the response.
