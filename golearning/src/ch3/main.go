package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

func main(){
	var a=1
	b:=3
	fmt.Printf("%T %T \n",a,b)

	cupArch:=runtime.GOARCH
	intSize:=strconv.IntSize

	fmt.Printf(cupArch,intSize,"\n")

	var f1 float32
	var f2 float64
	fmt.Printf("%f %f \n",f1,f2)


	var bt byte=2
	var ru rune='中'
	fmt.Printf("%T %T \n",bt,ru)

	var a1,a2=1,2
	var c int
	temp:= a1*a1+a2*a2
	fmt.Printf("%T \n",temp)

	c=int(math.Sqrt(float64(temp)))
	fmt.Printf("%T %d \n",c,c)

	p:=1
	ptr:=&p
	fmt.Printf("%T %T \n",p,ptr)
}

