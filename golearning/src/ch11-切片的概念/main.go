package main

import "fmt"

func changeEle(arr []int){
   arr[0]=1000
}

func main(){
   var arr [8] int =[8]int{0,1,2,3,4,5,6,7}

   s1:=arr[2:6]
   fmt.Println(s1)

   s2:=arr[:6]
   fmt.Println(s2)
   

   s3:=arr[2:]
   fmt.Println(s3)

   s4:=arr[:]
   fmt.Println(s4)
   
   changeEle(s4)
   fmt.Println(s4)
   fmt.Println(arr)



   var s0 []int
   s0=[] int{1,2,3}

   fmt.Printf("s0=%v,len=%d,cap=%d\n",s0,len(s0),cap(s0))

   s5:=make([]int ,5,5)
   fmt.Printf("s5=%v,len=%d,cap=%d\n",s5,len(s5),cap(s5))
}