package piscine

func MakeRange(min, max int) []int {
	size := 0
	Slice1 := make([]int, size)
	if max > min {
		size = max - min
		Slice1 = make([]int, size)
		for i := min; i < max; i++ {
			Slice1[i-min] = i
		}
	}
	return Slice1
}
