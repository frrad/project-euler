from fractions import Fraction 

start = Fraction(1,1)

top = 18
pool = set()
pool.add((start, 1))

for i in xrange(top):
  peel = set()
  for (frac1, size1) in pool:
    for (frac2, size2) in pool:
      size = size1 + size2
      if size > top: continue
      peel.add((frac1 + frac2, size))
      peel.add((1/(1/frac1 + 1/frac2), size))

  pool |= peel
  print len(pool)


