import euler

total = 0

for x in xrange(2, 2000000):
    if euler.primeQ(x):
        total += x


print total