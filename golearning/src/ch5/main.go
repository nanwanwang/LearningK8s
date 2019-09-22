package main

import "fmt"

func operatorDemo() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("a+b=%d\n", c)
	c = a - b
	fmt.Printf("a-b=%d\n", c)
	c = a * b
	fmt.Printf("a*b=%d\n", c)
	c = a / b
	fmt.Printf("a/b=%d\n", c)
	c = a % b
	fmt.Printf("a%%b=%d\n", c)
	a++
	fmt.Printf("a++=%d\n", a)
	a--
	fmt.Printf("a--=%d\n", a)
}

func relationDemo() {
	var a int = 21
	var b int = 10
	if a == b {
		fmt.Println("a==b")
	} else {
		fmt.Println("a!=b")
	}

	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a<=b")
	}
}

func logicDemo() {
	var a, b bool = true, false
	if a && b {
		fmt.Println("a和b都是真")
	} else {
		fmt.Println("a 和 b有假")
	}

	if a || b {
		fmt.Println("a或者b至少一个为真")
	} else {
		fmt.Println("a和b都为假")
	}

	if !a {
		fmt.Println("a为假")
	} else {
		fmt.Println("a为真")
	}
}
func main() {
	operatorDemo()
	relationDemo()
	logicDemo()
}
