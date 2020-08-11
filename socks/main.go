package main

import (
	"fmt"
	"os"
)

func socks(n int, arr []int) (matchingPairs int) {
	var result int

	counter := make( map[int]int )
	for _, row := range arr {
		counter[row]++
	}

	for _, v := range counter {
		result += v/2
	}

	return result
}

func main() {
	var socksTotal int

	_,err := fmt.Fscanf(os.Stdin, "%d", &socksTotal)
	if err != nil {
		fmt.Println(err)
	}

	arr := make([]int, socksTotal)
	y := make([]interface{}, len(arr))
	for i := range arr {
		y[i] = &arr[i]
	}

	total, err := fmt.Scanln(y...)
		if err != nil{
			fmt.Println(err)
		}

	arr = arr[:total]

	if total != socksTotal{
		fmt.Println("Color & total socks doesn't match")
		return
	}

	matchingSocks := socks(socksTotal, arr)

	println(matchingSocks)
}
