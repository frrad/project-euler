import euler

primes = []
n=1

#super slow (~1min). use a sieve 
while euler.prime(n) < 2000000:
    primes.append(euler.prime(n))
    n+=1

print sum(primes)