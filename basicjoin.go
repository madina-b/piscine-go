package piscine

func BasicJoin(strs []string) string {
	strMain := ""
	for i := range strs {
		strMain = strMain + strs[i]
	}
	return strMain
}
