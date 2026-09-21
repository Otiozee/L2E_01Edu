# Exercise 2: Role-Based Mentoring and Feedback

## the code I gave AI:
``` python
def is_palindrome(s):
    left = 0
    right = len(s) - 1
   
    while left < right:
      
        while left < right and not s[left].isalnum():
            left += 2
    
        while left < right and not s[right].isalnum():
            right -= 1
       
        if left < right:
            if s[left].lower() = s[right].lower():
               
                return left
            left -= 1
            right += 1
   
    return -1
```
## Prompt:
- Act as a mentor reviewing this code. Provide corrections and explain improvements.

## AI Response:
AI responded like a real tutor by stating the problems detected in the code as well as the cosequencies, as documented below:

1. Incrementing left by 2 instead of 1  
   Current: while left < right and not s[left].isalnum(): left += 2  
   Problem: This skips every second character while skipping non-alphanumeric symbols, so valid letters or digits get missed entirely.  
   Improvement: Change the increment to 1 so every character is checked properly.  
   Corrected line: while left < right and not s[left].isalnum(): left += 1  
   This ensures no alphanumeric characters are accidentally skipped.

2. Using assignment (=) instead of comparison (==)  
   Current: if s[left].lower() = s[right].lower():  
   Problem: A single = performs assignment, not equality check. This causes either a SyntaxError (in most modern Python versions inside if conditions) or completely broken logic.  
   Improvement: Replace = with == to correctly compare the lowercase versions of the characters.  
   Corrected line: if s[left].lower() == s[right].lower():  
   This fixes the core comparison so the function can actually detect matching characters.

3. Returning left (an index) instead of a boolean  
   Current: return left when characters match  
   Problem: The function name is_palindrome strongly suggests it should return True or False, but it returns a position (or -1 later). This confuses anyone using the function.  
   Improvement: Change the logic to return False as soon as a mismatch is found, and return True only after the entire string has been checked without mismatches.  
   Correct approach: return False on mismatch, and return True at the end of the loop if no mismatch occurred.

4. Moving pointers in the wrong direction after comparison  
   Current: left -= 1 and right += 1  
   Problem: After comparing two characters, the pointers move outward instead of inward. This re-examines already-checked positions and very often creates an infinite loop.  
   Improvement: After a successful match, move both pointers toward the center: left += 1 and right -= 1.  
   Corrected lines:  
   left += 1  
   right -= 1  
   This is the standard two-pointer movement for palindrome checking, it guarantees progress and eventual termination.

5. Case-insensitive comparison is broken  
   Current: Comparison never happens correctly due to the = typo in point 2.  
   Problem: Even if the typo were fixed in other ways, mixed-case palindromes like "Racecar" would fail without proper lowercasing.  
   Improvement: Always compare using .lower() on both sides (which is already attempted but broken by the typo). Once == is used, this will correctly ignore case differences.  
   No additional code is needed beyond fixing the comparison operator because the .lower() calls are already in the right place.

   ### AI Corrected version of the code: A Classic two-pointer
```python
def is_palindrome(s: str) -> bool:
    left = 0
    right = len(s) - 1
    
    while left < right:
        # Skip non-alphanumeric from left
        while left < right and not s[left].isalnum():
            left += 1
            
        # Skip non-alphanumeric from right
        while left < right and not s[right].isalnum():
            right -= 1
            
        # Now compare (case insensitive)
        if s[left].lower() != s[right].lower():
            return False
            
        left += 1
        right -= 1
        
    return True
```
   
## Reflection
The "act as a coding mentor" instruction meaningfully shaped the response in a positive way, for this context: It made the feedback warmer, more structured for learning, more interactive, and more focused on growth rather than just pointing out mistakes.