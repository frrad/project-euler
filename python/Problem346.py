import unittest


class Test1(unittest.TestCase):

    def test(self):
        self.assertEqual(
            strong_repunits(50), [1, 7, 13, 15, 21, 31, 40, 43])


class Test2(unittest.TestCase):

    def test(self):
        self.assertEqual(
            repunit_sum(1000), 15864)


def strong_repunits(under):
    repunits = set([1])
    base = 2
    while base + base ** 2 < under:
        repunit = base + base ** 2 + 1
        while repunit < under:
            repunits.add(repunit)
            repunit *= base
            repunit += 1
        base += 1
    return sorted(list(repunits))


def repunit_sum(top):
    return sum(strong_repunits(top))

# unittest.main()

print repunit_sum(10 ** 12)
