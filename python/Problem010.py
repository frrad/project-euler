import euler

total = 0
x = 1

euler.primeCache(2000000)

while euler.prime(x) < 2000000:
    total += euler.prime(x)
    x+=1

print total