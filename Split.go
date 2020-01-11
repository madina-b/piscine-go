package main

import "fmt"

func main() {
	Split("HelloHAhowHAareHAyou?", "HA"))
}
n:=0
func Split(str, charset string) [] string {
	charsetRune:=[] rune(charset)
	lenCharset:=0
	for i:=range charsetRune {
		lenCharset++
	}
	strRune:=[] rune(str)
	lenStr:=0
	for i:=range strRune {
		lenStr++
	}
	match:=make([]rune,lenCharset) 
	k:=0
	n:=0
	j:=0
	for i:=n; i<lenCharset; i++ { 
		for j:=j; j<lenStr; j++ {
			if charsetRune[i]==strRune[j] {
				match[k]=strRune[i]
				k++
				break
			}
			n:=j
			k=0
			break
			
		}
	}
}