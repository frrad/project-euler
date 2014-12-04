from math import sqrt as sqrt

epsilon = 10 ** -17


def next(a, b):
    return b, (-1 + 2 * b ** 2 + sqrt((b ** 2 - 1) * (a ** 2 - 1))) / b


# Returns True if too long.
def gauge(first_step, target_num):
    a, b = 0, first_step
    i = 0
    while i < target_num and b < 1:
        i += 1
        a, b = next(a, b)

    if b > 1:
        return True

    return False

# Binary search to find optimal first position


def search(blips):

    blips = 200

    a, b = .000001, 1

    while b - a > epsilon:
        c = .5 * (a + b)
        if gauge(a, blips) == gauge(c, blips):
            a, b = c, b
        else:
            a, b = a, c

    return c

top = 200

a, b = 0, search(top)
x_locations = [a, b]

for i in xrange(top):
    a, b = next(a, b)
    x_locations.append(b)


points = [(x, 0 if x > 1 else sqrt(1 - x ** 2)) for x in x_locations]


total = 0
for i, pt in enumerate(points[:-1]):
    total += (points[i + 1][0] - pt[0]) * pt[1]
print total * 4
