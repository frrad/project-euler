import euler

n=1

while euler.divisors(euler.triangle(n)) < 500:
    n += 1

print euler.triangle(n)