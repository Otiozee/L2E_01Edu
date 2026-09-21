# Part A – Honest self check

### How I've used AI for coding so far:  
Mostly copy-paste when I'm stuck after some trials. Sometimes I ask it to fix my bugs or explain errors. I also use it to write whole functions for leetcode-style problems.

### Do I ask before trying?  
Yeah… too often. If the problem looks long or has recursion I just throw it at ChatGPT right away.

### Can I explain code I submitted?  
Simple stuff yes - loops, lists, basic ifs. But anything with classes, dfs, dynamic programming - I usually just nod and hope they don't ask follow-ups.

### If AI disappears in exam/interview?  
I'd fail most medium/hard questions. I'd probably solve easy ones and panic on the rest.

### Current learner type:  
I'm like 55% Learner A right now. I treat it like a cheat sheet, but I’m learning on-the-go how to effectively use AI. I want to flip to mostly Learner B and use AI like a teacher who makes me think harder.

### Reflection:  
Right now I'm kind of lazy with AI, but I'm learning on the go. I get quick wins but I forget some fast and I panic when I can't ask. I want to be the guy who can think through problems even if the internet dies or AI is banned in the company. So I'm trying to force myself to struggle more first.

# Part B – Palindrome (The Right Way)

## Step 1 – My first attempt (no AI)

Pseudocode I thought of:  
```
INPUT string
SET string to lowercase
REMOVE all characters that are not letters or numbers
IF string is equal to reverse of string
    RETURN true
ELSE
    RETURN false
```

My code:

```python
def is_palindrome(s):
    # step 1: clean the string
    cleaned = ""
    for char in s:
        if char.isalnum():
            cleaned += char.lower()
    
    # step 2: compare with reverse
    return cleaned == cleaned[::-1]

print(is_palindrome("racecar"))          # True
print(is_palindrome("hello"))            # False
print(is_palindrome("A man a plan a canal Panama"))  # True
print(is_palindrome(""))                 # True? I think yes
print(is_palindrome(" "))                # True after clean
```

- Comments: I used isalnum() and lower because I learned it from my previous quest (Ethical-ai). Slicing is easy but maybe not best.

## Step 2 – Asked AI after (I pasted code + asked for complexity, edges, alternatives)
### What I learned:

- Time: O(n), space O(n) because new string + reversed string
- Edge cases I missed: very long string (but python ok), emojis?, non-english letters (isalnum may fail on accented chars)
- Better: two pointers without creating cleaned string - skip bad chars while comparing. That saves memory for huge inputs.
- I now feel I could write a string reverse function or remove duplicates without help.

### Reflection:
Struggling first made me actually read isalnum docs in my head. If I just asked for code I wouldn't know why two pointers is better. I built a little mental model: "extra space is often the price of clean code - sometimes worth paying, sometimes not".

# Part C – Variations (tried myself first)

```Python
def fancy_palindrome(s):
    left = 0
    right = len(s) - 1
    
    while left < right:
        # skip non-alnum from left
        while left < right and not s[left].isalnum():
            left += 1
        # skip from right
        while left < right and not s[right].isalnum():
            right -= 1
        
        if left < right:
            if s[left].lower() != s[right].lower():
                # return position in ORIGINAL string where mismatch happened
                return left
            left += 1
            right -= 1
    
    return -1  # is palindrome
```
    
- Tested: works on the panama phrase, "hello", "No 'x' in Nixon"

### Asked AI: "Did I miss edges? More efficient?"
- AI said: good job skipping, but unicode letters with accents might need extra handling (str.casefold() instead of lower), also very long string with many skips is still O(n). Looks solid.

# Part D – My Fairness Contract

### I will use AI when:

- I've tried at least 15-30 mins (or I'm really really stuck and frustrated)
- I want to know WHY my code is slow/wrong
- After I have something working and want to see smarter ways

### I will NOT use AI when:

- Timed coding interview or exam
- School homework that's meant to test me
- Learning basics like loops, functions, recursion for first time

### I know I'm using fairly when:

- I can turn off screen and still explain every line
- I can solve a similar problem next week without chat
- I feel smarter after, not just "done"

Signed,
zeotokpa
February 14, 2026

# Part E – Real scenarios

- Interview caching question: If I always asked "write LRU cache" I’ll probably be in a difficult position at that moment. Without understanding the trade-offs, like why hashmap + linkedlist is often the go-to, or what I'd sacrifice with other approaches. I’d sound like I’m reciting, not thinking.
- 2 AM prod bug: If code is AI-generated and I never understood it, I'd be googling at 3 AM or waking senior dev. Nightmare.
- New tech no AI knows yet: I'd have to actually read docs, run experiments, break things - if I never practiced that I'd be helpless.

### Reflection:
Using AI fairly now forces me to build real thinking muscle. When AI is gone or wrong or new stuff comes, I won't freeze because I already trained myself to reason first. Shortcuts today = weakness tomorrow. I want to be ready for real jobs, not just pass quests.

# Part F – Irreplaceable skills

### Ratings 1-5:

- Problem decomposition: 3/5
- Systems thinking: 2/5 - lowest
- Critical evaluation: 3/5
- Debugging mindset: 3/5
- Conceptual understanding (why): 2.5/5

- Lowest = Systems thinking
### 3 actions this week (no AI help):

- Every day I'll pick one small program and draw boxes+arrows on paper showing how data moves between functions/classes.
- After I finish any code, write 3 sentences: "If I to add something new, which files break and why?"
- Read one short article about how DNS works (no code) and explain it out loud to my mirror like I'm teaching others.

