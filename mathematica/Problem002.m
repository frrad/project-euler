Print[Select[TakeWhile[Fibonacci@Range[150], # < 4*10^6 &], EvenQ]//Total];
