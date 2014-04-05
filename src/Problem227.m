players = 100;
interest = players/2 +1;

d[in_] := Module[{a, b}, {a, b} = Sort[in];	 Min[b - a, a + players - b] ];

row[i_] := 
  Table[If[i == 1, 0, 
		   If[i == j, 1/2, 
			  If[d[{i, j}] == 1, 2/9, If[d[{i, j}] == 2, 1/36, 0]]]], {j, 1, players}];
mat = Table[row[i], {i, 1, players}];

start = {Table[z[i], {i, 1, players}]};
stat = {Table[z[i] + 1, {i, 1, players}]};

N[z[interest] /. (Transpose[start] == mat.Transpose[stat] //Solve)[[1]][[interest]], 10]
