package main

import "fmt"

// func main() {

// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	AdvancedSortWordArr(result, Compare)
// 	fmt.Println(result)
// }

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	k := 0
	for i := range array {
		k = i
	}
	k++
	for i := range array {
		for j := i + 1; j < k; j++ {
			if f(array[i], array[j]) >= 0 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
func Compare(a, b string) int {

	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1
}