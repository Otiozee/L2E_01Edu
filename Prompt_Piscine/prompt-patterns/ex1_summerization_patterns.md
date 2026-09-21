# Exercise 1: Summarization Pattern

## Weak Prompt:
“Summarize this text.”
- This tends to produce uneven results: sometimes too long, sometimes missing the point, or adding personal opinions.

## Refined Version:
You are a concise and objective summarizer.  
Summarize the following text in exactly 3 bullet points.  
Focus only on the central arguments, main findings or key decisions.  
Ignore examples, anecdotes, background stories, and minor details.  
Use neutral language. Start each bullet with a strong verb.

## Test with two different inputs: News article, and technical blog.

### Input 1: News article
Gunmen on motorcycles killed at least 50 people and abducted women and children during coordinated overnight attacks in Zamfara State, northwest Nigeria. A local lawmaker reported the rising death toll as rescue and search operations continue. The assaults highlight ongoing failure to contain banditry and jihadist groups despite increased military efforts. In response, Nigeria has deepened security ties with the United States, including the recent arrival of 100 US personnel for training missions and prior joint airstrikes requested by Nigerian authorities against Islamist targets.

Model output:
```markdown
- Reports high-casualty motorcycle attacks and mass abductions by armed groups in Zamfara State.
- Underscores persistent insecurity and challenges in containing banditry and jihadist activity.
- Notes strengthened US-Nigeria military cooperation through training deployments and coordinated operations.
```
### Input 2: Technical blog
The Central Bank of Nigeria released its 2026 Fintech Report, marking the first comprehensive sector review since the 2022 Payment Systems Vision. It calls for coordinated regulation, simplified licensing, and stronger institutional engagement. Stakeholders push for centralized regulatory channels, open banking infrastructure to enable secure data exchange, and cross-border passporting within Africa. The report also examines cryptocurrency potential for remittances and inclusion, recommending regulatory sandboxes for testing CBDCs, digital assets, and trade finance use cases while addressing compliance costs and infrastructure gaps.

Model output:
```markdown
- Presents first major fintech sector review since 2022, advocating coordinated regulation and easier licensing.
- Recommends centralized engagement channels, open banking frameworks, and African cross-border arrangements.
- Explores cryptocurrency and digital asset opportunities while proposing sandboxes for CBDC and trade finance innovation.
```
## Analysis
The refined version more explicitly follows the "news article and technical blog" suggestion while staying focused. The pattern still enforces tight, consistent 3-bullet point outputs. The output remained neutral, focused, no fluff, and they handles both hard news (violence and geopolitics), and industry/technical content (fintech regulation) with the same structure and reliability. The weak single-sentence prompt would again produce inconsistent, rambling, or opinionated results on these same texts.