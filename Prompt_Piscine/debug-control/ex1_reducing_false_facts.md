# Exercise 1: Reducing False or Invented Facts

## Weak Prompt That Produces Errors
"List the top 5 inventions created by humans on Mars."

Model response
```markdown
- Solar-powered oxygen farms

- Underground glass cities

- Martian gravity stabilizers

- Red dust energy converters

- Autonomous sandstorm shields
```
### Analysis
This prompt leads to made-up information because humans have not created inventions on Mars. Since the scenario itself is fictional, the model fills in completely fabricated answers. The output may sound realistic, but it is not factually correct at all.

## Improved Prompt That Restricts to Real Data
"List five real inventions developed on Earth in the 20th century and briefly describe their purpose."

Model response
```markdown
- Airplane – Enabled faster long-distance travel and global transportation.

- Television – Allowed transmission of visual media and entertainment.

- Antibiotics – Helped treat bacterial infections and saved millions of lives.

- Computer – Made data processing and automation possible.

- Internet – Connected people globally for communication and information sharing.
```
## Comparison and Differences
- The first prompt encourages fabrication because it asks for inventions on Mars, which do not exist. The model tries to generate a response which resulted in false information.
- The revised prompt limits the request to real, verifiable inventions from the 20th century. This makes the output fact-based and reliable.

The key difference is that the improved prompt:

 - Grounds the task in reality

 - Provides a clear time frame

 - Avoids fictional assumptions

 - Reduces the risk of hallucinated content

 ## Reflection
 This exercise shows that vague or unrealistic prompts often lead to inaccurate answers. By adding constraints and specifying real-world conditions, the quality and reliability of the output improves significantly.
