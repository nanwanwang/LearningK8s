package main

import "fmt"

func main(){
	var a int =58
	var b int = 10
	var c int 
	c= a+b
	fmt.Printf("a+b=%d\n",c)

	c=a-b
	fmt.Printf("a-b=%d\n",c)

	c=a*b
	fmt.Printf("a*b=%d\n",c)

	c=a/b
	fmt.Printf("a/b=%d\n",c)
	c=a%b
	fmt.Printf("a%%b=%d\n",c)
	a++
	fmt.Printf("a++=%d\n",a)
	a--
	fmt.Printf("a--=%d\n",a)

	c=a>>2
	fmt.Printf("a>>2=%d\n",c)


}