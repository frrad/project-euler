top = 10

memo = dict()

def win(a,b):
    A, B = a, b

    if a > b:
        return win(b, a)

    if (a,b) in memo:
        return memo[(a,b)]

    # assume a <= b

    # If already empty pile, we have lost
    if a == 0: return False

    while b > a:
        b-=a
        if not win(a, b):
            memo[(A, B)] = True
            return True


    memo[(A,B)] = False
    return False


for a in xrange(1,top):
    print ''.join(('x' if win(a,b) else ' ' for b in xrange(1,top)))

print sum( [x+y if not  win(x,y) else 0 for x in xrange(1, top+1) for y in xrange(x+1, top+1)])

