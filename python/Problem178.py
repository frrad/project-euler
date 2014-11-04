class path:
  def __init__(self, start=0, end=0, length=0):
    self.start = start
    self.end = end
    self.length = length
    self.seen = set()

  def key(self):
    pathstr = ''
    for i in xrange(10):
        pathstr += '1' if i in self.seen else '0'
    return '%d,%d,%d,%s' % (self.start, self.end, self.length, pathstr)

  def parse(self, instring):
    data = instring.split(',')
    self.start = int(data[0])
    self.end = int(data[1])
    self.length = int(data[2])
    self.seen = set()
    for i,s in enumerate(data[3]):
      if s == '1':
        self.seen.add(i)

  def pandigital(self):
    for i in range(10):
      if i not in self.seen:
        return False
    return True

  def next(self):
    out = []

    up = path()
    up.parse(self.key())
    up.end -= 1
    up.length += 1
    up.seen.add(up.end)

    if up.end >= 0: out.append(up.key())
    
    down = path()
    down.parse(self.key())
    down.end += 1
    down.length += 1
    down.seen.add(down.end)
    
    if down.end <= 9: out.append(down.key())

    return out

# Quantity of such numbers with exactly n digits
def total_digits(n):

  pool = dict()    
    
  for i in xrange(1,10):
    test = path(i,i,1)
    test.seen.add(i)
    pool[test.key()] = 1

  print pool

  for i in xrange(n - 1):
    peel = dict()

    for key in pool:
      joe = path()
      joe.parse(key)
      for follow in joe.next():
        peel[follow] = pool[key] + peel.get(follow, 0)

    pool = peel
    print i+2, len(pool)

  answer = 0

  for key in pool:
    road = path()
    road.parse(key)
    if road.pandigital():
      answer += pool[key]

  return answer

answer = sum(map(total_digits, range(5,41)))

print "\nAnswer: %d" % answer
