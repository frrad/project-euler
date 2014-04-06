p[k_, n_] := 1 - (Sum[Binomial[n, s] Binomial[s, k - s] (k)!/2^(k - s), {s, Ceiling[k/2], k}]/n^k);
expr = FullSimplify[p[k,n]];
k = 20000; n = 1000000;
N[expr, 10]
