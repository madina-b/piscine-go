//remove append function
package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

// func main() {
// 	PrintNbrBase(-922, "0123456789")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-12-5, "01")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "0123456789ABCDEF")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "choumi")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "a-")
// 	z01.PrintRune('\n')
// }

func PrintNbrBase(nbr int, base string) {
	baserune := []rune(base)

	result := []rune{}
	lenbase := 0
	lenresult := 0
	for _ = range baserune {
		lenbase++
	}
	div := lenbase
	if lenbase <= 1 {
		fmt.Print("NV")
		return
	}
	for i := 0; i <= lenbase-1; i++ {
		if baserune[i] == '+' || baserune[i] == '-' {
			fmt.Print("NV")
			return
		}
	}
	for i := 0; i <= lenbase-2; i++ {

		for j := i + 1; j <= lenbase-1; j++ {
			if baserune[i] == baserune[j] {
				fmt.Print("NV")
				return
			}
		}
	}
	if nbr < 0 {
		nbr = nbr * -1
		fmt.Print("-")
	}
	for i := nbr; i > 0; i = i / div {
		digit := baserune[i%div]
		result = append(result, digit) // chose str=digit+str
	}
	for _ = range result {
		lenresult++
	}
	for k := lenresult - 1; k >= 0; k-- {
		z01.PrintRune(result[k])
	}
}
