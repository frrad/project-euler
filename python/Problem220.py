def add(a,b):
    return (a[0] + b[0], a[1]+ b[1])

#scalar vector multiplication
def multiply(a, x):
    return (a * x[0] , a * x[1])

def L(a):
    return (-1*a[1], a[0])

def R(a):
    return (a[1], -1*a[0])

def dumb(n, steps):
    chart = word(n)
    x,v = (0,0), (0,1)
    counter = 0
    for step in chart:
        if step == 'F':
            counter += 1
            x = add(x,v)
        if step == 'R':
            v = R(v)
        if step == 'L':
            v = L(v)
        if counter == steps:
            return x
    
def word(n):
    if n ==0:
        return 'Fa'
    ans = ''
    for letter in word(n-1):
        if letter == 'a':
            ans += 'aRbFR'
        elif letter == 'b':
            ans += 'LFaLb'
        else:
            ans += letter
    return ans
        


def magic(n, steps):
    if n == 0:
        if steps == 0:
            return (0,0)
        else: 
            return (0,1)


    if steps % 2 == 0:
        if n%2 == 1:
            return multiply(2,  magic(n -1, steps/2))
        else:
            return  magic(n -1, steps/2)

    stoops = steps/2

    a, b = magic(n-1,stoops), magic(n-1, stoops + 1)
    if stoops % 2 == 1:
        a,b = b,a

    if (a[0] != b[0] and a[1] != b[1]): 
        if (b[0] > a[0]) != (b[1] < a[1]):
            return (a[0], b[1])
        else:
            return (b[0], a[1])

    a,b = multiply(2,a), multiply(2,b)


    if a[0] == b[0]:
        delta = (a[1] - b[1]) /2
        avg = (a[1] + b[1]) /2
        return (a[0] + delta, avg)

    if a[1] == b[1]:
        delta = (a[0] - b[0]) / 2
        avg = (a[0] + b[0]) / 2
        return (avg, a[1] - delta)



def magicWrap(k, n):
    point = magic(k,n)

    a = point[0]
    b = point[1]
    modder = 8
    if k%modder == 0:
        return(a,b)
    if k%modder == 1:
#        if n%2 ==0:
        return (a/2 + b/2,-a/2 + b/2)

    if k%modder == 2:
        return (-b,a)
    if k%modder == 3:
        return (a/2-b/2,a/2+b/2)
        

    return (0,0)

def sketch(alist):
    a = reduce(min, (i[0] for i in alist))
    b = reduce(max, (i[0] for i in alist))
    c = reduce(min, (i[1] for i in alist))
    d = reduce(max, (i[1] for i in alist))


    for i in range(a,b+1):
        print "".join(["X" if (i,j) in alist else " " for j in range(c, d+1)])

k = 17
print 'k=', k

# fold  = [dumb(k,i)      for i in range(1,1 + 2**k)]  
# foool = [magicWrap(k,i) for i in range(1 + 2**k)]  

# sketch(foool)
# print '========='
# sketch(fold)
# print '========='


scoop = 0.
scope = 40.
for i in range(40):
    if magicWrap(k,i) != dumb(k,i):
        print "%7s:%7s:%7s (%d)" % (dumb(k,i), magic(k,i), magicWrap(k,i),i)
        scoop += 1
        
print scoop/scope*100
