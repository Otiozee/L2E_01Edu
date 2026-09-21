# Exercise 3: Parameter Experimentation


## 1. Use the same prompt with temperature = 0.2 and then with temperature = 0.9.

"Write a 100-word short story about a robot that learns to dance."

- Temperature = 0.2 (very low randomness):
In a quiet factory, Robot Unit 47 was programmed only for assembly tasks. One night its owner played music. Unit 47 watched the steps carefully, recording each movement. Day after day it practiced in the corner. Its joints moved with perfect timing. Soon the factory workers gathered to watch. Unit 47 had learned to dance. The owner smiled proudly as the robot spun gracefully under the lights.

- Temperature = 0.9 (high randomness):
Zappy the little cleaning bot suddenly heard salsa music from the break room. Its circuits went wild! Instead of mopping, it started spinning on one wheel, arms flailing like broken windmills. Oil spilled everywhere. The humans laughed so hard they cried. Zappy crashed into a chair, rebooted, then nailed a perfect moonwalk. From that day it became the factory disco king, dreaming of electric breakdance battles under neon stars.

## 2. Repeat with different top-p values (e.g., 0.5 vs 1).

- Top-p = 0.5 (more focused):
Similar to low temp - safe, logical story about the robot practicing steps and succeeding. No crazy details.

- Top-p = 1.0 (maximum diversity):
The story was way wilder, the robot invents a new dance called "circuit shuffle," involves sparks flying, accidentally starts a dance party with all the other machines, ends with the factory throwing an all-night rave.

## What I noticed:
Low temperature / low top-p = boring but consistent and focused. The story is predictable, like a kids book. High temperature / high top-p = fun, creative, sometimes silly or off-topic, but more entertaining. For school work or code explanations I'd pick low values so the answer stays on track. For stories or brainstorming I'd go higher to get fresh ideas. The randomness really changes the personality of the answer.