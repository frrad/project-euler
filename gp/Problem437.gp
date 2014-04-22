top = 100000000

default(primelimit,top) 

check1(x,rt)= znorder((1+rt)/2)==x-1
check2(x,rt)= znorder((1-rt)/2)==x-1

total=5

{
forprime( p=1, top, 
  	if(Mod(p,5)==1 || Mod(p,5)==4,
  	 	rt = sqrt(Mod(5,p));
  	 	if(check1(p,rt)==1 || check2(p,rt)==1,
	 		total += p
  	 	)
 	)
)
}

print(total)
quit