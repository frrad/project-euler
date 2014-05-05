import euler

def d(n): return euler.divisorSigma(n,1)-n
def amicableQ(a): return d(a)!=a and d(d(a)) == a

print sum((n for n in xrange(2,10000) if amicableQ(n)))
