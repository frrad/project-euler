import euler

top = 1000
pals = [a*b for a in range(top) for b in range(top) if euler.palindrome(a*b)]
print max(pals)