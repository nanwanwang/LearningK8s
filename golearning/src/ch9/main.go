package main

import "fmt"

var c int =15

func main(){

  var a,b,c int
  a=12
  b=11
  c=24
  fmt.Println(a,b,c)

  var p =10


  funcValREF(p)

  fmt.Println(p)

  funcPtr(&p)

  fmt.Println(p)

  fmt.Println(&p)

  var add =&p



  fmt.Println(*&p)
  fmt.Println(add)
  fmt.Println(*&add)

  var add2=&add
  fmt.Println(*&add2)
}

func funcValREF(a int) {
  a= 1000
}

func  funcPtr(a *int)  {
  *a=1000
}
