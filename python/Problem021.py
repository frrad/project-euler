import euler

def d(n):
    return euler.divisorSigma(n,1)-n

def amicableQ(a):
    b = d(a)
    if b==a:
        return False
    if d(b) == a:
        return True
    return False

print sum([n for n in range(2,10000) if amicableQ(n)])