import euler

top = 14

# All single digit numbers are Harshad
niven = [[(x, x) for x in xrange(1, 10)]]


# Given a (right truncatable) Harshad, how can we add a digit while
# remaining right truncatable Harshad?
def expand(number, total):
    answer = []
    base = number * 10
    for x in xrange(10):
        if (base + x) % (total + x) == 0:
            answer.append((base + x, total + x))
    return answer


# Given a right truncatable Harshad, which primes truncate to it?
def primespand(number):
    answer = []
    for x in [1, 3, 7, 9]:
        consider = (number * 10) + x
        if euler.primeQ(consider):
            answer.append(consider)
    return answer

# expand the table of right truncatable Harshad numbers to the appropriate
# size.
for i in xrange(top - 2):
    next = []
    for entry in niven[-1]:
        next.extend(expand(*entry))
    niven.append(next)

# Check which of our candidates are strong.
strong_harshad = []
for size in niven:
    for (number, total) in size:
        if euler.primeQ(number / total):
            strong_harshad.append(number)


# Finally take the strong, right truncatable Harshad numbers and find
# their corresponding primes.
answer = 0
for harsh in strong_harshad:
    answer += sum(primespand(harsh))

print answer
