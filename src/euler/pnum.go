package euler

import (
	"fmt"
)

type PNum struct {
	primes  map[uint64]uint64
	primelist [][2]uint64
}

//returns pointer to pnum equivalent to n
func IntPNum(n int) *PNum{
	response := new(PNum)
	response.primes = make(map[uint64]uint64)

	factors := Factors(int64(n))
	for _, x := range factors {
		prime, exponent :=  x[0], uint64(x[1])
		response.set( uint64( PrimeN(prime) ),  exponent)
	}

	return response
}

//returns pointer to pnum equivalent to n
func UInt64PNum(n uint64) *PNum{
	response := new(PNum)
	response.primes = make(map[uint64]uint64)

	factors := Factors(int64(n)) //of course this is silly
	for _, x := range factors {
		prime, exponent :=  x[0], uint64(x[1])
		response.set( uint64( PrimeN(prime) ),  exponent)
	}

	return response
}



//returns n! as a PNum
func FacPNum(n uint64) *PNum{

	response := new(PNum)
	response.primes = make(map[uint64]uint64)

	for primeIndex := uint64(1); primeIndex <= uint64 (PrimePi( int64(n)) ); primeIndex++ {
		exp := uint64 (0)
		for primePow := uint64(1); UInt64Exp( uint64 (Prime( int64 (primeIndex) )), primePow) <= n; primePow++ {
			exp += n / (UInt64Exp(uint64 ( Prime( int64(primeIndex))) , primePow))
		}
		response.set(primeIndex, exp )
	}

	return response
}



//returns the uint64 value 
func (p *PNum) UInt64() (val uint64) {
	val = uint64(1)
	for prime, exp := range p.primes {
		val *=   UInt64Exp( uint64( Prime( int64(prime) ) ) ,  exp)  
	}
	return
}


//returns the nth divisor 
func (p *PNum) Divisor(n uint64) (d *PNum){
	d = new(PNum)
	d.primes = make(map[uint64]uint64)

	if len(p.primelist) == 0{
		p.populatePrimelist()
	}

	for i := 0; i < len(p.primelist) && n > 0; i++ {
		//fmt.Println(p.primelist[i][0],p.primelist[i][1])
		d.set(p.primelist[i][0], n%(p.primelist[i][1]+1) )
		n /= p.primelist[i][1]+1
	}

	return
}

//returns a/b assuming b divides a
func Quotient(a, b *PNum) *PNum {
	quot := new(PNum)
	quot.primes = make(map[uint64]uint64)

	for p, exp := range a.primes {
		quot.set(p, exp - b.primes[p])
	}

	return quot
}


func (p *PNum) Mod( n uint64) uint64 {
	nold := n
	ntemp := UInt64PNum(n)
	flag := true

	reduce := uint64(1)

	ptemp := make(map[uint64]uint64)
	for prime, exp := range p.primes{
		ptemp[prime] = exp
	}

	for prime, exp := range ntemp.primes {

		//not divisible
		if ptemp[prime] < exp{
			flag = false
		}

		if ptemp[prime] !=0 {
			
			if ptemp[prime] > exp {
				reduce *= UInt64Exp(uint64( Prime(int64(prime)) ), exp)
				ptemp[prime] -=  ntemp.primes[prime]
				delete(ntemp.primes , prime) 			
			}else{
				reduce *= UInt64Exp( uint64( Prime(int64(prime)) ), ptemp[prime])
				ntemp.primes[prime] -=  ptemp[prime]
				delete(ptemp  , prime)
			}

		}

	}

	if flag {
		return 0
	}

	n = ntemp.UInt64()

	fmt.Println(ptemp , "over", n, "reduced by ", reduce)




	for prime, exp := range ptemp {
		if  UInt64Prime(prime) > n{


			modulo := UInt64Prime(prime) % n

			factors := Factors(int64(modulo)) //of course this is silly

			for _, x := range factors {
				aprime, anexponent :=  uint64(x[0]), uint64(x[1])
				ptemp[aprime] += anexponent * exp
			}

			delete(ptemp, prime)

		}

		//Can do further tricks: what if prime^exp > n ?
	}

	fmt.Println(ptemp)

	for prime, exp := range ptemp{
		reduce *= UInt64Exp( UInt64Prime(prime), exp)
		reduce %= nold
	}

	return reduce

}

func (p *PNum) Divisors() uint64 {
	div := uint64(1)
	for _, i := range p.primes{
		div *= i+1
	}
	return div
}

func (p *PNum) set(a,b uint64) {
	p.primes[a] = b
}

func (p *PNum) populatePrimelist() {
	temp := make([]uint64,0)

	for index := range p.primes {
		temp = append(temp, index)
	}

	//fmt.Println(temp)
	SortUInt64(temp)


	for _, index := range temp {
		p.primelist = append(p.primelist, [2]uint64{index, p.primes[index]})
	}
}