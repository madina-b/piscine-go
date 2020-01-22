package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

// func main() {

// 	PrintNbrBase(919617, "01")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(753639, "CHOUMIisDAcat!")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-174336, "CHOUMIisDAcat!")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-661165, "1")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-861737, "Zone01Zone01")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "0123456789ABCDEF")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "choumi")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "-ab")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-9223372036854775808, "0123456789")
// 	z01.PrintRune('\n')

	// 11100000100001000001
	// HIDAHI
	// -MssiD
	// NV
	// NV
	// 7D
	// -uoi
	// NV
	// -9223372036854775808
}

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

	upperBound := 9223372036854775807
	lowerBound := -9223372036854775807

	if nbr > upperBound || nbr < lowerBound {
		if base == "0123456789" {
			fmt.Print(nbr)
			return
		}
	}
	if nbr < 0 {
		nbr = nbr * -1
		fmt.Print("-")
	}
	for i := nbr; i > 0; i = i / div {
		digit := baserune[i%div]
		result = append(result, digit)
	}
	for _ = range result {
		lenresult++
	}
	for k := lenresult - 1; k >= 0; k-- {
		z01.PrintRune(result[k])
	}
}
