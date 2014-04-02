def push(frac):
    frac = (frac[0]*10, frac[1])
    frac = (frac[0]%frac[1] ,frac[1])
    return frac

def findLength(frac):
    seen = set()
    old = dict()

    i=1

    while not frac in seen and frac[0] != 0:
        seen.add(frac)
        old[frac] = i

        i += 1
        frac = push(frac)

    if frac in seen:
        return  i-old[frac]


longest,best = 1,-1
for denom in xrange(1,1000):
    if findLength((1,denom)) > longest:
        best, longest = denom, findLength((1,denom))


print best
