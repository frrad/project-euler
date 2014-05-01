n = 1000; k = n/2 - 1;

from[B_, A_] :=
Module[{aHalf, bHalf, a, b},
   {a, b} = {A, B};
   aHalf = If[a > k, True, False]; 
   bHalf = If[b > k, True, False];
   If[aHalf, a -= k + 1;];
   If[bHalf, b = b - k - 1;];
   0 +
   If[a == -1 || b == -1, If[a == b, 1, 0], 0] +
   If[a >= 0 && a + 1 == b && False == aHalf == bHalf, (n - 2 a - 2)/n, 0] +
   If[a + 1 == b && True == aHalf == bHalf, (n - 2 a - 2)/n, 0] +
   If[a == b && aHalf == False && bHalf == True, 1/n, 0] +
   If[a == b && aHalf == bHalf, a/n + 1/n, 0] +
   If[b == -1 && aHalf, a/n + 1/n, 0] +
   If[a != -1 && b == -1 && ! aHalf, a/n, 0]
];

IminusQ = SparseArray[{i_, j_} /; Abs[i - j] == n/2 || Abs[i - j] <= 1 :> 
    If[i == j, 1, 0] - from[j - 1, i - 1], {n, n}];

N[Total[Inverse[ IminusQ][[1]]], 10]

