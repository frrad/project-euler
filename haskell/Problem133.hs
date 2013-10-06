import Euler

divides p = 1 `elem` trials
	where
		trials = zipWith mod (map (10^) (stop pows)) (repeat p)
		pows = map (gcd (p-1)) (map (10^) [0..])
		stop (x1:x2:xs) = if x1 == x2 then [x1] else x1:stop(x2:xs)

answer = sum $ takeWhile (<100000) $ filter (not.divides) (map prime [1..])
main = putStrLn(show(3+answer))