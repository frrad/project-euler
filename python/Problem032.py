def panDigits(a,b):
    pan = str(a*b) + str(a) + str(b)
    data = sorted(list(pan))
    return "".join(data) == '123456789'

def digs(a,b):
    x1, x2 = 10**(a-1), 10**(a)
    y1, y2 = 10**(b-1), 10**(b)
    z=[a*b for a in range(x1,x2) for b in range(y1,y2) if panDigits(a,b)]
    return sum(set(z))

print digs(1,4) + digs(2,3)
