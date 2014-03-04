collatz = dict({1:1})

def chainlength(n):
    if n in collatz:
        return collatz[n]
    if n%2 ==0:
        collatz[n] = 1 + chainlength(n/2)
        return chainlength(n)
    collatz[n] =  1 + chainlength(3*n +1)
    return chainlength(n)

best, size = -1, -1

for x in xrange(1,1000000):
    if chainlength(x) > size:
        # print "{}\t{}".format(x, chainlength(x))
        best, size =x, chainlength(x)

print best