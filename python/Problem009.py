perimeter = 1000

possible = [(a,b,perimeter-a-b) for a in range(1,perimeter) for b in range(a,(perimeter-a)/2)]
for (a,b,c) in possible:
    if a**2 + b**2 == c**2:
        print a*b*c