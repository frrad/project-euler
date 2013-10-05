palindrome :: String -> Bool
palindrome [] = True
palindrome [_] = True
palindrome xs = (head xs) == (last xs) && palindrome (init (tail xs))

macks = maximum [x * y | x <- [100..999], y <- [100..999] , palindrome (show (x*y))]

main = putStrLn (show(macks))