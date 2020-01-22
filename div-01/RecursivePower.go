package main

import (
	"fmt"
)

//func main() {

	arg1 := -7
	arg2 := -2
	fmt.Println(RecursivePower(arg1, arg2))
	arg1 = -8
	arg2 = -7
	fmt.Println(RecursivePower(arg1, arg2))
	arg1 = 4
	arg2 = 8
	fmt.Println(RecursivePower(arg1, arg2))
	arg1 = 1
	arg2 = 3
	fmt.Println(RecursivePower(arg1, arg2))
	arg1 = -1
	arg2 = 1
	fmt.Println(RecursivePower(arg1, arg2))
	arg1 = -6
	arg2 = 5
	fmt.Println(RecursivePower(arg1, arg2))

}

// 0
// 0
// 65536
// 1
// -1
// -7776

func RecursivePower(nb int, power int) int {
	a := 1
	if power < 0 || power > 32 {
		return 0
	} else if nb == 1 {
		return 1
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power-1)
	}
	return a
}
