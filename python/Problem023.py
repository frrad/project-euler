import euler

def isAbundant(n):
    pdsum = euler.divisorSigma(n,1) - n
    return pdsum > n

top = 28123
abundant = []

for x in range(10,top):
    if isAbundant(x): abundant.append(x)

doubles = set(range(top))

for j, number in enumerate(abundant):
    i = j
    while abundant[i] + number <= top:
        doubles.discard(number + abundant[i])
        #print number + abundant[i]
        i += 1

print sum(doubles)
