package piscine

func UltimateDivMod(a *int, b *int) {
	a1 := a / b
	b1 := a % b
	*a = a1
	*b = b1
}
