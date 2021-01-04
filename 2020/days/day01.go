package days

import (
	inputs "../inputs"
	"fmt"
)

func One() {
	var slice = inputs.Day1

	for i, n := range slice {
		for j, m := range slice {
			if i >= j {
				continue
			}
			if m+n == 2020 {
				fmt.Printf("1st return: %d\t %d \t\t ==> %d \n", m, n, m*n)
			}

			for k, o := range slice {
				if i >= j || i >= k || j >= k {
					continue
				}
				if m+n+o == 2020 {
					fmt.Printf("2nd return: %d\t %d \t %d \t ==> %d \n", m, n, o, m*n*o)
				}
			}
		}
	}
}
