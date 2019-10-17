package piscine

func Swap(a *int, b *int) {
	a1 := *a
	b1 := *b
	*a = b1
	*b = a1

}
