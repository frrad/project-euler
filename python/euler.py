import math
###################
###Combinatorics###
###################

#'pascal's triangle' implementation of binomial function
#probabley not very efficient
binomCache = dict()
def choose(n, k):
    if n<0 or k<0 or k>n:
        return 0
    if n==0 or k==0:
        return 1
    if (n,k) in binomCache:
        return binomCache[(n,k)]
    answer = choose(n-1, k) + choose(n-1, k-1)
    binomCache[(n,k)] = answer
    return answer


####################################
###Primality, Factorization, etc.###
####################################
primes = [2]
def prime(n):
    if len(primes) >= n:
        return primes[n-1]
    while len(primes) < n:
        start = primes[-1] + 1
        while not primeQ(start):
            start += 1
        primes.append(start)
    return prime(n)


def primeCache(n):
    list = [True for x in range(n)]
    list[0] = False
    i = 0

    while i < int(math.ceil(math.sqrt(n))):
        j = i + 1
        while not list[j]:
            j += 1
        j += 1

        count = j 
        while j - 1 < len(list) - count:
            list[count + j - 1] = False
            j += count
        i += 1
        while not list[i]:
            i += 1

    global primes 
    primes = []
    for j in xrange(0, n):
        if list[j]:
            primes.append(j+1)


#speed up using prime list?
def factor(n):
    if n <= 2:
        return [n]
    test = 2
    while test**2 <= n:
        if n%test == 0:
            sub = factor(n/test)
            sub.append(test)
            return sub
        test += 1
    return [n]

def factors(n):
    basic = factor(n)
    basic.reverse() #want ascending order
    left, right = 0,0
    answer = []
    while left < len(basic):
        while right < len(basic) and basic[left] == basic[right]:
            right += 1
        answer.append((basic[left], right-left))
        left = right
    return answer


def primeQ(n):
    if n < 2: return False
    d = 2
    while d * d  <= n:
        if n % d == 0:
            return False
        d += 1
    return True


#Euclid's algorithm
def GCD(a,b):
    while b !=0:
        t = b
        b = a%b
        a = t  
    return a

def LCM(a,b):
    return a*b / GCD(a,b)

#number of divisors (\sigma_0)
def divisors(n):
    fac = factors(n)
    d = 1
    for (_, a) in fac:
        d *= a+1
    return d

#number of divisors (\sigma_0)
def divisorSigma(n,k):
    if n==1: return 1
    if k==0: return divisors(n)
    fac = factors(n)
    d = 1
    for (p, a) in fac:
        d *= (p**(k * a +k) - 1) / (p**k - 1)
    return d



#######################
###Special Sequences###
#######################
fibCache = dict({0:1, 1:1})
def fib(n):
    if n in fibCache:
        return fibCache[n]
    fibCache[n] = fib(n-1) + fib(n-2)
    return fib(n)

def palindrome(n):
    #make sure we've got string
    word = str(n)

    for i in range(len(word)/2):
        if word[i] != word[-1-i]:
            return False

    return True

factorialCache = dict({1:1, 2:2})
def factorial(n):
    if n < 1: return 1 
    if n in factorialCache:
        return factorialCache[n]
    factorialCache[n]=factorial(n-1)*n
    return factorial(n)

def triangle(n):
    return n*(n+1)/2

##############
###ANALYSIS###
##############

def lagrange(points):
    points = map(lambda list: map(float,list), points)
    n = len(points)

    def ans(x):
        ans = 0
        for j in range (n):
            p = points[j][1]
            for k in range (n):
                if k == j: continue
                p *= (x - points[k][0])/(points[j][0] - points[k][0])
            
            ans += p
        return ans
    
    return ans
