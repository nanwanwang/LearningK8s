package main

import "fmt"

func main() {
	a := 10
	if a < 20 {
		fmt.Println("a<20")
	} else if a == 20 {
		fmt.Println("a=20")
	} else {
		fmt.Println("a>20")
	}

	switch {
	case a > 10:
		fmt.Println("a>10")
	case a == 10:
		fmt.Println("a=10")
	case a < 10:
		fmt.Println("a<10")
	default:
		fmt.Println("未知")
	}

	switch a {
	case 10:
		fmt.Println("a=10")
		fallthrough
	case 11,13,14:
		fmt.Println("a=11")
	case 12:
		fmt.Println("a=12")
	default:
		fmt.Println("未知")
	}
}
