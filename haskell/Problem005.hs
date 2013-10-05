nDivide x n = length (takeWhile (==0) (zipWith mod (repeat x) [1..n])) == n

ans = head [ x | x <- [2520, 5040..] , nDivide x 20] 

main = putStrLn (show(ans))