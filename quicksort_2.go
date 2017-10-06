package main

import "fmt"

func quickSort(parr []int) []int {
	piv := parr[0]
	larr := make([]int, 0)
	rarr := make([]int, 0)
	retarr := make([]int, 0)
	l := len(parr)

	for i := 1; i < l; i++ {
		if parr[i] < piv {
			larr = append(larr, parr[i])
		} else {
			rarr = append(rarr, parr[i])
		}
	}

	if len(larr) > 1 {
		larr = quickSort(larr)
		print(larr)
	}

	if len(rarr) > 1 {
		rarr = quickSort(rarr)
		print(rarr)
	}

	retarr = append(retarr, larr...)
	retarr = append(retarr, piv)
	retarr = append(retarr, rarr...)

	return retarr
}

func print(arr []int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	print(quickSort(arr))
}
