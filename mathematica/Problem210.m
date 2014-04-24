n = 10^9;
omega = n/4;
flour[x_] := If[IntegerQ[x], Max[0, x - 1], Floor[x]]
a = ParallelSum[flour[Sqrt[i omega - i^2]], {i, 0, omega}];
b = ParallelSum[flour[(1/2) (1 + Sqrt[2 omega (2 i + 1) - (2 i + 1)^2])], {i, 0, omega - 1}];
(a + b)*2 + 24 omega^2
