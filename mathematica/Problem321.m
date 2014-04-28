(*Reduce[n^2+2n\[Equal]k(k+1)/2&&n>0&&k>0,Integers]*)
a = Table[Simplify[1/8 (-8 - (3 - 2 Sqrt[2])^x (-4 + Sqrt[2]) + (4 + Sqrt[2]) (3 + 2 Sqrt[2])^x)],{x, 1, 40}];
b = Table[Simplify[1/8 (-8 + (3 - 2 Sqrt[2])^x (4 + Sqrt[2]) - (-4 + Sqrt[2]) (3 + 2 Sqrt[2])^x)],{x, 1, 40}];
Sort[Union[a, b]][[1 ;; 40]] // Total
