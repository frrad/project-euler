import euler

def facDigits(x): return sum(map(euler.factorial,map(int, str(x))))

print sum([x for x in range(3,10**5) if x == facDigits(x)])
