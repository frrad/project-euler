# scalar vector multiplication
def multiply(a, x):
    return (a * x[0], a * x[1])

memo = dict({(0, 1): (0, 1), (0, 0): (0, 0)})


def magic(n, steps):
    if (n, steps) in memo:
        return memo[(n, steps)]

    if steps % 2 == 0:
        if n % 2 == 1:
            memo[(n, steps)] = multiply(2, magic(n - 1, steps / 2))
        else:
            memo[(n, steps)] = magic(n - 1, steps / 2)
        return magic(n, steps)

    steps /= 2

    a, b = magic(n - 1, steps), magic(n - 1, steps + 1)
    if steps % 2 == 1:
        a, b = b, a

    if (a[0] != b[0] and a[1] != b[1]):
        if (b[0] > a[0]) != (b[1] < a[1]):
            return (a[0], b[1])
        else:
            return (b[0], a[1])

    a, b = multiply(2, a), multiply(2, b)

    if a[0] == b[0]:
        delta = (a[1] - b[1]) / 2
        avg = (a[1] + b[1]) / 2
        return (a[0] + delta, avg)

    delta = (a[0] - b[0]) / 2
    avg = (a[0] + b[0]) / 2
    return (avg, a[1] - delta)


def magicWrap(k, n):
    (a, b) = magic(k, n)

    if k % 8 == 0:
        return(a, b)
    if k % 8 == 1:
        if n % 2 == 0:
            return (a / 2 + b / 2, -a / 2 + b / 2)
        if n % 2 == 1:
            return ((a + b) / 2, b / 2 - a / 2)
    if k % 8 == 2:
        return (b, -a)
    if k % 8 == 3:
        return ((-a + b) / 2, (-a - b) / 2)
    if k % 8 == 4:
        return (-a, -b)
    if k % 8 == 5:
        return ((-a - b) / 2, (a - b) / 2)
    if k % 8 == 6:
        return (-b, a)
    if k % 8 == 7:
        return ((a - b) / 2, (a + b) / 2)

    return None


print magicWrap(50, 10 ** 12)
