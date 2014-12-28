# Number of triangles
def Triangle(n):
    return ((n * n / 6) + 1) / 2 - (n / 4) * ((n + 2) / 4)


def seive(top):
    seive_data = [0 for i in xrange(top)]

    accumulate = 0

    for n in xrange(2, top):
        # if prime
        if seive_data[n] == 0:

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


print seive(1 + 10 ** 7)
