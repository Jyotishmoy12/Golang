package main

import "fmt"

func main() {

	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < len(a); i++ {
		var key int = a[i]
		var j int = i - 1

		for j >= 0 && a[j] > key {

			a[j+1] = a[j]
			j = j - 1
		}

		a[j+1] = key

	}

	fmt.Println(a)
}
