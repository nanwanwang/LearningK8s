package main

import "fmt"
var c1,c2=6,"golang"

var ( 
	a1=7
	a2="demon"
)

func initValue(){
	var a,b int = 2,4
	fmt.Println(a,b)
	var s string = "string"
	fmt.Println(s)
	var s1,s2,s3=5,"hello",true
	fmt.Println(s1,s2,s3)

	s4,s5,s6:=5,"world",false
	fmt.Println(s4,s5,s6)
	fmt.Println(c1)
	fmt.Println(a2)
}


func  main()  {
  var a int
  var b string
  a=3
  b="567"
  fmt.Println(a,b)
  fmt.Printf("%d %q",a,b)
  initValue()
}
