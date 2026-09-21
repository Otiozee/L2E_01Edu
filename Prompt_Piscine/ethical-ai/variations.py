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

print(is_palindrome("A man a plan a canal Panama"))  # should be -1
print(is_palindrome("hello"))                        # should return 0
print(is_palindrome("Able was I ere I saw Elba"))    # -1
