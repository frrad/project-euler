(*Assuming[k > 0, TransformedDistribution[(k a + 1)^2, a \[Distributed] UniformDistribution[{0, 1}]] // PDF]*)
one[k_, x_] := If[x > 1 && x < 1 + 2 k + k^2, 1/(2 Sqrt[x] k), 0];
two = Convolve[one[k, x], one[k, x], x, y];
Print["Computed PDF."];
ans = Assuming[Element[k, Integers], Integrate[two, {y, (k - 1/2)^2, (k + 1/2)^2}]] // FullSimplify;
Print["Computed Expected Value. Summing..."];
N[Sum[k ans /. k -> j, {j, 1, 10^5}], 11]

