package main

import "fmt"

func main() {

println("-------------------------")
	a := 10
LOOP:
	for a < 20 {
		a++
		if a == 15 {
			goto LOOP
		}
		fmt.Println(a)
	}
}
