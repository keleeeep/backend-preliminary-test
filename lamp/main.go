package main

import (
	"fmt"
	"math"
)

func main(){
	var trips float64 = 100
	lamp := int(math.Sqrt(trips))
	fmt.Println(lamp)
}
