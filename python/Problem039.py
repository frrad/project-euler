def solutions(perimeter):
    return filter(isRight, ((a,b,perimeter-a-b) for a in xrange(1,perimeter/3) for b in xrange(a, (perimeter-a)/2)))

def isRight((a,b,c)):
    return a**2+b**2==c**2

print max(xrange(1001), key=lambda x : len(solutions(x)))
