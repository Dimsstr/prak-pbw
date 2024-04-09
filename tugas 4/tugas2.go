package main

import "fmt"

func main() {
	fmt.Println("Bubble Short")
	var arrayNumber [20]int
	for i := 0; i < len(arrayNumber); i++ {
		arrayNumber[i] = len(arrayNumber) - i
	}

	i := 0
	for i < len(arrayNumber) {
		if i != (len(arrayNumber) - 1) {
			if arrayNumber[i] > arrayNumber[i+1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i+1]
				arrayNumber[i+1] = x
				i--
			}
		}
		if i > 0 {
			if arrayNumber[i] < arrayNumber[i-1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i-1]
				arrayNumber[i-1] = x
				i--
			}
		}
		i++
	}

	fmt.Println("Setelah dilakukan pengurutan.")
	fmt.Println(arrayNumber)
}
