from euler import choose

colors_in_rainbow = 7
total_balls = 70
balls_to_draw = 20
balls_per_color = total_balls / colors_in_rainbow


# number of partitions of n into exactly k parts
def partition(n, k):
    if n == k and n == 0:
        return 1
    if n <= 0 or k <= 0:
        return 0
    return partition(n - k, k) + partition(n - 1, k - 1)


# how many configurations contain n or fewer distinct colors
def or_fewer(n):

    # which colors to avoid
    answer = choose(colors_in_rainbow, n)

    # how many such configurations
    answer *= choose(balls_per_color * n, balls_to_draw)

    return answer


def exactly(n):
    return or_fewer(n) - or_fewer(n - 1)


print sum((n * exactly(n) for n in xrange(2, colors_in_rainbow + 1))), choose(total_balls, balls_to_draw)
