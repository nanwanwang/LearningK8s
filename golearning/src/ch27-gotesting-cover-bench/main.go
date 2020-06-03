package main

import (
	calculate "ch27-gotesting-conver-bench/calculates"
	"fmt"
)
func main(){
	for i:=0;i<=100;i++{
		printf, err := fmt.Printf("%d id event?%v\n", i, calculate.Even(i))
	}
}