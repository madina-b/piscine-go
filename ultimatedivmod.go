package piscine

func UltimateDivMod(a *int, b *int) {
	a1 := *a
	b1 := *b
	*a = a1 / b1
	*b = a1 % b1
}
