// Package pkg simply adds the number
package pkg

import "fmt"

// Sum prints the sum of numbers
func Sum(n ...int) {
	sum := 0
	for _, v := range n {
		sum += v
	}
	fmt.Println(sum)
}
