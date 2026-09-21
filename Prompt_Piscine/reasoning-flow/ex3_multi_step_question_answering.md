# Exercise 3: Multi-Step Question Answering

## Multi-step question:
What is 30% of the difference between 150 and 90?

### Prompt 1: Restate the problem
Model response:
```markdown
We need to find 30% of the result when 90 is subtracted from 150.
```
### Prompt 2: Break it into smaller steps
Model response:
```markdown
Step 1: Find the difference between 150 and 90.

Step 2: Calculate 30% of that result.
```
### Prompt 3: Solve each step before giving the final answer
Model response:
```markdown
Step 1: Find the difference.
150 − 90 = 60

Step 2: Find 30% of 60.
30% means 0.30 in decimal form.
0.30 × 60 = 18
```
## Review the Reasoning
- The subtraction is correct: 150 − 90 = 60.

- 30% converted to decimal form is 0.30.

- 0.30 × 60 = 18.

- Each step follows logically and calculations are accurate.

- Final answer is 18

## Analysis:
The reasoning process clearly restates the problem, separates it into logical steps (first finding the difference, then calculating the percentage), and shows how each intermediate result, like obtaining 60 from 150 − 90. This directly supports the next calculation, leading to the final answer of 18.