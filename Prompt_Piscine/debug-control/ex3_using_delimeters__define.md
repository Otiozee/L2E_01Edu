# Exercise 3: Using Delimiters to Define Context

## Prompt Without Delimiters
"Read the passage below and answer the question.

Lagos is the largest city in Nigeria and serves as the country’s main commercial center. Although Abuja is the capital city, Lagos remains the economic hub, with major industries including finance, entertainment, and transportation. The city is located in the southwestern part of Nigeria along the Atlantic coast.

Question:
What industries are mentioned as major industries in Lagos?"

### Model response
```markdown
The major industries in Lagos include finance, entertainment, transportation, and other commercial activities.
```
### Analysis
The answer is mostly correct, but it slightly adds “other commercial activities,” which was not directly stated in the passage. Without strict instructions, the model may expand beyond the exact wording of the text.

## Prompt With Delimiters and Clear Restrictions
"Answer the question using only the text inside the delimiters. Do not add any outside information.

Lagos is the largest city in Nigeria and serves as the country’s main commercial center. Although Abuja is the capital city, Lagos remains the economic hub, with major industries including finance, entertainment, and transportation. The city is located in the southwestern part of Nigeria along the Atlantic coast.

Question:
What industries are mentioned as major industries in Lagos?"

### Model response
```markdown
Finance, entertainment, and transportation.
```
## Comparing outputs
- In the first version, the response slightly expanded beyond the exact wording of the passage. In the second version, the delimiters and clear instruction forced the answer to stay strictly within the provided text.

- This shows that adding delimiters and explicit constraints improves accuracy and reduces the chance of unnecessary additions.