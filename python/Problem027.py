import euler

top = 1000

def lth(a,b):
    n = 0
    while euler.primeQ(n**2 + a*n + b):
        n+=1
    return n

max = 0

for a in xrange(-1*top+1, top,2): #odds only
    b = 1
    while euler.prime(b)<top:
        p = euler.prime(b)
        if lth(a,p) > max:
            max, ans = lth(a,p), a*p
        b+=1

print ans
