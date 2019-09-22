package main

import (
	"fmt"
	"math"
)
const (
	s1 = "golang"
	MAX=10
)

func enumDemo(){
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday, Saturday)
}

func typeDefAlias()  {
	type MyInt1 int
	type MyInt2 = int
	var  i int = 1
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	println(i,i1,i2)
}

func  main()  {
	const  s string  = "Hello"
	const a,b  =3,4
	fmt.Println(s,a,b)
	var c int
	c= int(math.Sqrt(a*a+b*b))
	fmt.Println(c)
	fmt.Println(s1,MAX)
	enumDemo()
	typeDefAlias()
}
