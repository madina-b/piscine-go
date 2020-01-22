package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// func main() {
// 	n := len(os.Args)
// 	if n < 4 {
// 		fmt.Println("Not enough arguments")
// 		os.Exit(1)
// 	}
// 	count, files := Counting(os.Args[1:])
// 	point := 0
// 	Name := len(files) > 1
// 	for j, f := range files {
// 		filename, err := os.Open(f)
// 		if err != nil {
// 			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
// 			os.Exit(0)
// 		}
// 		if Name == true {
// 			fmt.Printf("==> %s <==\n", f)
// 			point = 1
// 		}
// 		read := make([]byte, int(count))
// 		_, er := filename.ReadAt(read, fileSize(filename)-int64(count))
// 		if er != nil {
// 			fmt.Println(er.Error())
// 		}
// 		for _, word := range read {
// 			z01.PrintRune(rune(word))
// 		}
// 		if j < len(files)-1-point {
// 			z01.PrintRune('\n')
// 		}
// 		filename.Close()
// 	}
// }
func fileSize(name *os.File) int64 {
	filename, err := name.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return filename.Size()
}
func Counting(args []string) (int, []string) {
	n := len(args)
	count := 0
	arrstr := []string{}
	for i, word := range args {
		var err error
		_, err = strconv.Atoi(word)
		if word == "-c" {
			if i >= n-1 {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
				os.Exit(1)
			}
			arg := args[i+1]
			count, err = strconv.Atoi(arg)
			if err != nil {
				arg = "'" + arg + "'"
				fmt.Printf("tail: invalid number of bytes: %s\n", arg)
				os.Exit(0)
			}
			continue
		}
		if err != nil {
			arrstr = append(arrstr, word)
		}
	}
	return count, arrstr
}