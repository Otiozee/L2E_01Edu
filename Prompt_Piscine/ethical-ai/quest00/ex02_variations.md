My attempt before asking:
```python
    def is_palindrome(s):
    # clean version: keep only letters/numbers, lowercase
    clean = ""
    for char in s:
        if char.isalnum():
            clean += char.lower()
    
    # check
    left = 0
    right = len(clean) - 1
    
    while left < right:
        if clean[left] != clean[right]:
            # return the position where it fails (0-based from start)
            return left, right
        left += 1
        right -= 1
    
    return -1  # means it is a palindrome

## tests
print(is_palindrome("A man a plan a canal Panama"))  # return -1
print(is_palindrome("hello"))                        # return 0
print(is_palindrome("Able was I ere I saw Elba"))    # return -1
```
- What I think I did ok: cleaning + two pointers + returning mismatch position

## What I asked AI: “Did I miss anything? Can it be more efficient?”

### AI pointed out:

- isalnum() is good but only ascii - if unicode emojis or accented letters come, might break (but fine for now)

- Building clean string still uses O(n) space - could do two pointers on original string and skip non-alnum chars

- Returning left is ok but maybe clearer to return the index in original string (debatable). Could short-circuit early on mismatch instead of building whole clean string

- I didn’t think about skipping without cleaning - that’s actually smarter for very long inputs.

## Reflection:
AI caught that I was still using extra memory even after switching to two pointers. The skip-non-alnum version is cleaner in memory. I learned something I wouldn’t have noticed alone.
