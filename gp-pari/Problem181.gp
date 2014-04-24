\\generating function code due to Michael Somos (OEIS)
{
T(n, k) = 
if( n<0 || k<0,
	0, 
	polcoeff(
		polcoeff( 
			prod(i=1, n, prod(j=0, i, 1 / (1 - x^i * y^j), 1 + x * O(x^n))
		), n)
	, k)
)
} 

print(T(40+60,40))
quit