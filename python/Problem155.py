from fractions import Fraction 

memo = dict()
memo[1] = set([Fraction(1,1)])

def caps(n):
  if n in memo:
    return memo[n]
  answer = set()
  for a in xrange(1,n):
    for x in caps(a):
      for y in caps(n-a):
        answer.add(x+y)
        answer.add(1/(1/x+1/y))
  memo[n]=answer
  return caps(n)


aggregate = set()
print 'i    new    total '
print '==  =====  ======='
for i in range(1, 18 + 1):
  this = caps(i)
  aggregate |= this
  print "%2d %7d %7d" % (i, len(caps(i)), len(aggregate))
