import euler

consider = [euler.fib(n) for n in range(70) if euler.fib(n) < 4000000 and euler.fib(n)%2 ==0]
print sum(consider)