from euler import choose

colors_in_rainbow = 7

balls_to_draw = 20
balls_per_color = 10
total_balls = balls_per_color * colors_in_rainbow


def only_red_and_orange(n):
    if n == 2:
        return 1

    return choose(n * balls_per_color, balls_to_draw) - sum([choose(n, x) * only_red_and_orange(x) for x in xrange(2, n)])


def only_n_distinct_colors(n):
    return choose(colors_in_rainbow, n) * only_red_and_orange(n)


print float(sum([a * only_n_distinct_colors(a) for a in range(2, 8)])) / choose(total_balls, balls_to_draw)
