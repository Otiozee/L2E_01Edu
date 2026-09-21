## My first try (before asking anyone):
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

```python
    def is_palindrome(s):
    # reverse the string and compare
    reversed_s = s[::-1]
    if s == reversed_s:
        return True
    else:
        return False

## tests
print(is_palindrome("racecar"))     # True
print(is_palindrome("hello"))       # False
print(is_palindrome("A man a plan a canal Panama"))  # False - because of spaces and case
```
## Comments:
I used slicing because it’s short and I remember it from somewhere. Probably not the most efficient for super long strings but it’s fine for now.

## After asking AI (I pasted my code and asked: what’s time complexity, edge cases, better ways?):

### I Learned:

- Time complexity: O(n), and it basically measures how many operations we perform.

- Space complexity: O(n) - because slicing creates a new string. O(n) here basically means We stored n things.

- Edge cases I missed: empty string should be True, single letter True, all same letters True, strings with numbers, very long string (but Python handles it ok)

- Better way: two pointers - start left=0, right=len-1, compare while moving in, no extra space used (O(1) space) which stores only 2 inters and touches n characters.

- I could now write a reverse-string function easily - just loop from end to start and build a new one, or use [::-1] still.

## Reflection:
- Doing it myself first made me realize how lazy the reverse trick is, it works but wastes memory.
- After seeing the two-pointer version I actually understand why interviewers like it.
- I feel like I could write either version now without looking.