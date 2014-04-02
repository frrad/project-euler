def dfive(n):
    return sum([int(d)**5 for d in str(n)]) == n

top = 6*9**5
print sum([n for n in xrange(2,top) if dfive(n)])
