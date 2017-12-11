import Data.Function (fix)

combine f a b = (f a) + 4*a + (f b) + b

minCost 1 = 0
minCost x = minimum (splitCosts minCost x)
splitCosts g x = map (\a ->(combine g a (x-a)))  [1..x-1]


split g x a = combine g (a+1) (x-a-1)
costF x
  | x<= 1 = 0
  | otherwise = split (\z -> cost!!z) x (div (x-2) 4)
cost = map costF [0..]
