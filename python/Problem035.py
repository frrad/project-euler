import euler

def rotate(x):
    return map(int,(str(x)[i:]+str(x)[:i] for i in range(len(str(x)))))
def circular(p): return all(map(euler.primeQ, rotate(p)))
def odd(p): return all([int(d)%2 == 1 for d in str(p)])
def ordered(p): return all(int(str(p)[0]) <= int(d) for d in str(p))

euler.primeCache(10**6)
cans = filter(lambda x : odd(x) and ordered(x) and circular(x), euler.primes)
answer = set.union({2}, *(rotate(x) for x in cans))

print len(answer)



