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
    if n ==0:
        if steps == 0:
            return (0,0)
        else: 
            return (0,1)

    if steps%2 ==0:
        return multiply(2,  magic(n -1, steps/2))
    
    m = n/2

    a, b = magic(m,steps-1), magic( m + 1, steps - 1)

    a,b = multiply(2,a), multiply(2,b)

    if a[0] != b[0] and a[1] != b[1]:
        if m % 2 == 0:
            return (b[0], a[1])
        else:
            return (a[0], b[1])

    if a[0] == b[0]:
        delta = (a[1] - b[1]) / 2
        avg = (a[1] + b[1]) /2 
        if m % 2 == 0:
            return (a[0] + delta, avg)
        else:
            return (a[0] - delta, avg)

    if a[1] == b[1]:
        delta = (a[0] - b[0]) / 2
        avg = (a[0] + b[0]) / 2
        if m % 2 == 0:
            return (avg, a[1] + delta)
        else:
            return (avg, a[1] - delta)


print dumb(10, 500)
print magic(10, 500)
