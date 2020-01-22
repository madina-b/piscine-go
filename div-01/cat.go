
package main

import (
    	"io/ioutil"
    	"os"

    	"github.com/01-edu/z01"
    )

    // func main() {
    // 	argument := os.Args
    // 	k := 0
    // 	for i := range argument {
    // 		k = i
    // 	}
    // 	if k == 0 {
    // 		end := -1
    // 		buffertmp := []byte{0}
    // 		i := 0
    // 		next := false

    // 		for end < 0 {

    // 			os.Stdin.Read(buffertmp)

    // 			s := string(buffertmp)
    // 			if i == 0 && s == "^" {
    // 				next = true
    // 			}
    // 			if i == 1 && next == true && s == "C" {
    // 				return
    // 			}
    // 			if i == 1 && next == true && s != "C" {
    // 				next = false
    // 				z01.PrintRune('^')
    // 			}
    // 			i++
    // 			if s == "\n" {
    // 				i = 0
    // 			}
    // 			for _, word := range s {
    // 				if next != true {
    // 					z01.PrintRune(word)
    // 				}
    // 			}
    // 		}
    // 	}
    	for i := 1; i <= k; i++ {
    		file, err := ioutil.ReadFile(argument[i])
    		if err != nil {
    			errarr := []rune(string(err.Error()))
    			for _, word := range errarr {
    				z01.PrintRune(word)
    			}
    			z01.PrintRune(10)
    			return
    		}
    		arr := []rune(string(file))
    		for _, word := range arr {
    			z01.PrintRune(word)
    		}
    	}
    }
