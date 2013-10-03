sum35 0 = 0
sum35 n 
	|n `mod` 3 == 0 ||  n `mod` 5 == 0 = n+sum35 (n-1)
	|otherwise = sum35 (n-1)
	
main = putStrLn (show(sum35 999))