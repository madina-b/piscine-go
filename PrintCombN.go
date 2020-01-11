package main

import "fmt"

//func main() {
	PrintCombN(2)

//}

func PrintCombN(n int) {
	for i := 0; i <= 10-n; i++ {
		fmt.Print(i)
		if n > 1 {
			n = n - 1
			PrintCombN(n)
		}
	}
}
