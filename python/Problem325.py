top = 10 ** 16

count_memo = dict()
weight_memo = dict()
stack = dict()


x = (1, 1)
while x[0] <= top:
    stack[x[0]] = x[1]
    x = (x[0] + x[1], x[0] + 2 * x[1])

stock = sorted(stack.keys())


# return the largest special fibonacci number <= x
def flour(x):
    ans = 0
    for key in stock:
        if key <= x:
            ans = key
        else:
            return ans
    return ans


def count(n):
    if n < 2:
        return 0

    if n in count_memo:
        return count_memo[n]

    floor_n = flour(n)
    remainder = n - floor_n

    if floor_n == n:
        ans = stack[n] - n + count(n - 1)
        count_memo[n] = ans
        return ans

    ans = (count(floor_n) + count(remainder) +
           (stack[floor_n] - floor_n) * remainder)
    count_memo[n] = ans
    return ans


def weight(n):
    if n < 2:
        return 0

    if n in weight_memo:
        return weight_memo[n]

    floor_n = flour(n)
    remainder = n - floor_n

    if floor_n == n:
        ans = n * (stack[n] - n - 1) + weight(n - 1) + \
            (1 - n + stack[n]) * (n + stack[n]) / 2
        weight_memo[n] = ans
        return ans

    ans = (weight(floor_n) +
           # Swath
           (floor_n - n) * (floor_n - stack[floor_n]) *
           (floor_n + 2 * n + stack[floor_n] + 3) / 2 +
           # Backtrack
           weight(n - floor_n) + count(n - floor_n) *
           (floor_n + stack[floor_n]))
    weight_memo[n] = ans
    return ans


def break_point(tip):
    def truth(a):
        return a * a * 5 < (tip * 2 - a) ** 2

    a, b = 0, tip
    while b - a > 1:
        c = (b + a) / 2
        if truth(a) != truth(c):
            a, b = a, c
        elif truth(b) != truth(c):
            a, b = c, b
    return a

breakpoint = break_point(top)

ans = (weight(breakpoint) + (breakpoint - top) *
       ((breakpoint + 1) ** 2 - top ** 2) / 2)
print ans, ans % 7 ** 10
