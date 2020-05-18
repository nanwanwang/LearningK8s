
package main


import "fmt"

func main(){

   var a,b,c int
   a=10
   b=20
   c=a+b
   fmt.Println(a,b,c)

   var p=20
   var ip *int
   ip=&p
   fmt.Println(ip)
   fmt.Println(*ip)

   *ip=30
   fmt.Println(p)
}