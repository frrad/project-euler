import euler

a = euler.lagrange(enumerate([1,3,13]))
b = euler.lagrange(enumerate([1,9,25]))
c = euler.lagrange(enumerate([1,7,21]))
d = euler.lagrange(enumerate([1,5,17]))

def layer(n): return (a(n), b(n), c(n), d(n))
def layers(n): return int(round(sum(layer(n))))
    
adjust = euler.lagrange([[1,1],[3,2],[5,3]])

print sum([layers(n) for n in range(int(adjust(1001)))])-3
