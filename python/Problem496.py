from math import sqrt as sqrt


def listerine(n):
    return ((a, b, (a ** 2 - b ** 2) / b) for a in xrange(n + 1) for b in xrange((a + 2) / 2, int((sqrt(1 + 4 * a ** 2)) - 1) / 2) if ((a * a) % b == 0))


def f(n):
    return sum((a[0] for a in listerine(n)))


def explore(b, limit=9999999):
    return list(filter(lambda x: True if x[1] == b and x[0] <= limit else False, listerine(5 * b)))


def boil(n):
    ans = 1
    d = 2
    while n != 1:
        if n % d == 0:
            n /= d
            ans *= d
            if n % d == 0:
                n /= d
            d -= 1
        d += 1
    return ans


def smartplore(b):
    boiled = steam[b]

    start = int(sqrt(b * b + b) + 1)

    end = 2 * b - 1

    a, b = (int(1 + start / boiled),  int(end / boiled))
    if a <= b:
        return - (-1 + a - b) * (a + b) * boiled / 2
    else:
        return 0

def smartplore2(b, end):
    boiled = steam[b]

    start = int(sqrt(b * b + b) + 1)


    a, b = (int(1 + start / boiled),  int(end / boiled))
    if a <= b:
        return - (-1 + a - b) * (a + b) * boiled / 2
    else:
        return 0


def efff(a):
    A = [sum(z[0] for z in explore(b)) for b in xrange((a + 2) / 2)]
    B = [sum(z[0] for z in explore(b, a))
         for b in xrange((a + 2) / 2, int((sqrt(1 + 4 * a ** 2)) - 1) / 2)]
    return sum(A) + sum(B)


def EFF(a):
    A = [smartplore(b) for b in xrange(3, (a + 2) / 2)]
    B = [smartplore2(b, a)
         for b in xrange((a + 2) / 2, int((sqrt(1 + 4 * a ** 2)) - 1) / 2)]
    return sum(A) + sum(B)


def steamer(top):
    x = 2
    while x < top:
        if steam[x] == 1:
            n = x
            while n < top:
                m = n
                while m < top:
                    steam[m] *= x
                    m += n
                n *= x * x
        x += 1


asdf = 10 ** 6
steam = list(1 for a in xrange(asdf))
steamer(asdf)

print EFF(asdf)
