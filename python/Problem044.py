import euler

top = 1
sums = []

while len(sums) == 0:
    pentagons = set((euler.pentagon(i) for i in xrange(1,top)))
    sums = filter(lambda (a,b): a+b in pentagons and abs(a-b) in pentagons , ((a,b) for a in pentagons for b in pentagons))
    #print sums, top
    top *= 2

print(abs(sums[0][0]-sums[0][1]))


