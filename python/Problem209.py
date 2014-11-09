def relate(ins):
    [a,b,c,d,e,f] = ins
    return [b,c,d,e,f,a^(b&c)]

def binary(n):
    ans = []
    while len(ans) < 6:
        ans = [n%2] + ans
        n/=2
    return ans

def unary(place):
    ans = 0
    for pea in place:
        ans *=2
        ans += pea
    return ans

def nextSet(n):
    return unary(relate(binary(n)))

def colorings(n):


    a = 0
    b = 1
    for i in xrange(n-1):
        a, b = a + b, a
    ans = a


    a = 1
    b = 0
    for i in xrange(n-1):
        a, b = a + b, a


    return ans + a + b

def cycles():
    cycles = []
    seen = set()
    for n in xrange(64):
        i = 0 
        while n not in seen:
            seen.add(n)
            i += 1
            n = nextSet(n)

        if i > 0:
            cycles.append(i)

    return cycles

print reduce(lambda a,b: a*b, map( colorings, cycles() ))

