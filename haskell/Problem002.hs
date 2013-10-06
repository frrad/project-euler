fib 0 = 1
fib 1 = 1
fib n = fibL !! (n-1) + fibL !! (n-2)

fibL = map fib [0..]

answer = sum . filter even $ takeWhile (<=4000000) fibL

main = putStrLn (show(answer))