package main

import (
	"fmt"
	"strings"
)

func main(){
	s := strings.Split("1.345.679", "")

	for i, v := range s {
		if v == "." {
			s = append(s[:i], s[i+1:]...)
		}
	}

	strLength := len(s)-1

	for i := range s {
		fmt.Println(s[i] + strings.Repeat("0", strLength))
		strLength--
	}
}


