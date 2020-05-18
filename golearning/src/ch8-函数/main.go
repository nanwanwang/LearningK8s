package main

import (
	//  "errors"
	"fmt"
	"math"
	 "os"
)

func operate(a,b int,op string) (int,error) {
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a-b,nil
	case "*":
		return a * b,nil
	case "/":
		return a/b,nil
	default:
	    return 0,fmt.Errorf("not support operate : %s ",op)


	}
}

func swap1(a,b int) (int,int){
	return b,a
}

func swap2(a,b int)(x,y int){
	x,y=b,a
	return
}

func compute(op func(int,int) int,a,b int) int{
	return op(a,b)
}

func main() {

   if c,err:=operate(2,3,"o"); err!=nil{
	   fmt.Println(err.Error())
   }else{
	   fmt.Println(c)
   }
  
   a,b:=swap1(3,4)
   fmt.Println(a,b)
   x,y:=swap2(7,8)
   fmt.Println(x,y)

   
   if file,err := os.Open("abc.txt"); err!= nil{
	   fmt.Println("打开文件有误")
   }else{
	   fmt.Println(file)
   }

   fmt.Println(compute(powInt,3,4))

   say(englishGreeting,"world")

   say(chinaGreeting,"世界")

   fmt.Println(sum(1,2,3,4,5))
}

func powInt(a,b int) int{
	return int(math.Pow(float64(a),float64(b)))
}

func say(g greeting,world string){
	fmt.Println(g(world))
}

type greeting func(name string) string

func englishGreeting(name string) string{
	return "Hello "+ name
}

func chinaGreeting(name string) string{
	return "你好"+name
}


func sum(nums...int) int{
	s:=0
  for i:=0;i<len(nums);i++{
	  s+=nums[i]
  }
  return s
}