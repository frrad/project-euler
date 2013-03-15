
package main

import "fmt"

func main() {
count:=0

	for twobucks := 0; twobucks <= 1; twobucks++ {

		for abuck := 0; abuck <= 2; abuck++ {

			for fiddy := 0; fiddy <= 4; fiddy++ {

				for twenty := 0; twenty <= 10; twenty++ {

					for ten := 0; ten <= 20; ten++ {

						for five := 0; five <= 40; five++ {

							for two := 0; two <= 100; two++ {

								for one := 0; one <= 200; one++ {

if twobucks*200 + abuck*100 + fiddy*50 + twenty*20 + ten*10 + five*5 + two*2 + one == 200 {count++}

								}
							}
						}

					}

				}
			}
		}
	}

	fmt.Println(count)
}

