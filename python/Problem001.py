five = set()
for n in xrange(0,1000,5):
    five.add(n)

three = set()
for n in xrange(0,1000,3):
    three.add(n)

consider = five.union(three)

sum = 0

while len(consider) > 0:
    sum += consider.pop()

print sum