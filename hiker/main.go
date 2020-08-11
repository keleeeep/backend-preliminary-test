package main

import (
	"fmt"
	"os"
)

func main() {
	var n, path, valleys int
	var step string

	fmt.Fscanf(os.Stdin, "%d", &n)

	var steps []string
	for i := 0; i < n; i++ {
		fmt.Fscanf(os.Stdin, "%s", &step)
		steps = append(steps, step)
	}

	for _, v := range steps {
		switch v {
		case "U":
			path += 1
			break
		case "D":
			path -= 1
			break
		}
	}

	if path == 0 {
		valleys = 1
	} else if path < 0{
		valleys = 0
	}

	fmt.Println(valleys)
}