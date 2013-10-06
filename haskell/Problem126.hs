import Data.Map (toList, fromListWith)

layerP a b c n = let m = n-1 in 4*m^2 + 4*(a+b+c-1)*m + 2*(a*b+a*c+b*c) 

cs k = toList $ fromListWith (+) [(c, 1) | c <- table]
	where 
		table =  concat trai 
		trai = [takeWhile (<= k) $ map (layerP a b c) [1..]| (a,b,c)<- admissable]
	 	admissable = [(a,b,c)| a<-[1..k], b<-(takeWhile (\x -> first a x x <= k) [a..]), c<-(takeWhile (\x -> first a b x <= k) [b..])]
	 	first a b c = 2*(a*b+a*c+b*c) 

answer = fst $ head[(a,b)| (a,b)<-cs 20000, b==1000]

main = putStrLn(show(answer))