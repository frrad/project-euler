module Euler
( 
  primeQ,
  prime
)where

primeQ 1 = False
primeQ n = if 0 `elem` trial then False else True
       where trial = map (mod n) (takeWhile (\x -> x^2 <= n) primes)

prime 1 = 2
prime n = head $ filter primeQ [(primes!!(n-2)) + 1 ..]

primes = map prime [1..]
