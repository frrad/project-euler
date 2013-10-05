factor :: Integral a => a -> [a]
factor 1 = []
factor x = p :factor (x `div` p)
	where p = head [div | div <- [2..] , x `mod` div == 0]

biggestFactor :: Integral a => a -> a
biggestFactor x = last (factor x)

main = putStrLn (show(biggestFactor 600851475143))