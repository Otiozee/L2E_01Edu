def is_palindrome(s):
    # reverse the string and compare
    reversed_s = s[::-1]
    if s == reversed_s:
        return True
    else:
        return False

# tests
print(is_palindrome("racecar"))     # True
print(is_palindrome("hello"))       # False
print(is_palindrome("A man a plan a canal Panama"))  # False -because of spaces and case