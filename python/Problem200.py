import euler

def twohundred(x):
    return '200' in [str(x)[i:i+3] for i in range(len(str(x))-2)]

def primeProof(x):
    xstr = str(x)
    for i in range(len(xstr)):
        for a in range(10):
            ystr = xstr[:i]+str(a)+xstr[i+1:]
            if euler.primeQ(int(ystr)): return False
    return True

euler.primeCache(200000)

squbes = [2*2*2*q*q for q in euler.primes]
stopper = max(squbes)
for p in xrange(1, len(euler.primes)):
    q = 2
    while euler.prime(p)**2 * euler.prime(q)**3 < stopper:
        if p == q : q +=1
        squbes.append(euler.prime(p)**2 * euler.prime(q)**3)
        q+=1

squbes.sort()
squbes = filter(twohundred, squbes)

i = 1
for can  in squbes:
    if primeProof(can):
        print can, i
        if i == 200:
            print can
            break
        i+=1
