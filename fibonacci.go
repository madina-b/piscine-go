package piscine

func Fibonacci(index int) int {
	a := 0
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {

		a = Fibonacci(index-2) + Fibonacci(index-1)
	}
	return a
}
