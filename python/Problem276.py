from euler import primeQ as primeQ

# Number of triangles


def Triangle(n):
    return ((n * n / 6) + 1) / 2


# Number of primitive triangles
def Primitive(n):
    if primeQ(n):
        return Triangle(n)

    ans = 0
    for i in xrange(2, n):
        if n % i == 0:
            ans += Primitive(i)

    return Triangle(n) - ans


def brute(n):
    return sum((Primitive(i) for i in xrange(2, n)))


def seive(top):
    seive_data = [0 for i in xrange(top)]

    accumulate = 0

    for n in xrange(2, top):
        # if prime
        if seive_data[n] == 0:
            print n
            seive_data[n] = Triangle(n)
        else:
            seive_data[n] = Triangle(n) - seive_data[n]

        delta = seive_data[n]
        accumulate += delta

        k = n + n
        while k < top:
            seive_data[k] += delta
            k += n

    return accumulate


print seive(10**7)


