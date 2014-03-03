def d(n):
    divisors = [d for d in range(1,n) if n%d ==0]
    return sum(divisors)

def amicableQ(a):
    b = d(a)
    if b==a:
        return False
    if d(b) == a:
        return True
    return False

print sum([n for n in range(1,10000) if amicableQ(n)])