package main

import (
	calculate "ch24-gotesting/calculates"
	"fmt"
)
func main(){
	for i:=0;i<=100;i++{
		fmt.Printf("%d id event?%v\n",i,calculate.Even(i))
	}
}