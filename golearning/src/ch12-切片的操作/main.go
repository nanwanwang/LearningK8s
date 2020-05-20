package main

import "fmt"

func main(){
	var s1 [] int;
	s1= append(s1,1)
	s1=append(s1,10)
	fmt.Println(s1==nil)
	fmt.Printf("s1=%v,len=%d,cap=%d\n",s1,len(s1),cap(s1))


	arr:=[...] int {0,1,2,3,4,5,6,7,8,9}
	s2:=arr[2:6]
	fmt.Println(s2,len(s2),cap(s2))


	fmt.Println("*****************")

	fmt.Println(arr)
	s2=append(s2,10)
	fmt.Println(s2,len(s2),cap(s2),arr)

	s3:=make([]int,1023)
	s3=append(s3,1)
	fmt.Println(cap(s3))

	s4:=[]int{1,2,3,4}
	s5:=make([]int,6,6)

	copy(s5,s4)
	fmt.Println(s5)

	s6:=[]int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println(s6,len(s6),cap(s6))
	s6=append(s6[:4],s6[5:]...)
	
	fmt.Println(s6,len(s6),cap(s6))

}