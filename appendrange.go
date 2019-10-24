package piscine

func AppendRange(min, max int) []int {
	var Slice1 []int
	if min < max {
		for i := min; i < max; i++ {
			Slice1 = append(Slice1, i)
		}
	}
	return Slice1
}
