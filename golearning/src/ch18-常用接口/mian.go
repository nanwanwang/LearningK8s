package main
import (
	"fmt"
	"sort"
)


type Appender interface{
	Append(int)
}

type Lener interface{
  Len() int
}

type List [] int

func (l List) Len() int{
	return len(l)
}

func (l *List) Append(val int){
	*l=append(*l,val)
}

func Counter(a Appender,start,end int){
	for i:=start;i<=end;i++{
		a.Append(i)
	}
}

func IsLarge(l Lener) bool{
	return l.Len()>8
}



type Course struct{
	Title string
	SubTitle string
}

func (c *Course)String()string{
	return fmt.Sprintf("[Course]{Title=%s,SubTitle=%s}",c.Title,c.SubTitle)
}

func main(){
//   var lst List
//   Counter(lst,1,10)

 // fmt.Println(IsLarge(lst))

//  使用指针  接收者为指针的可以调用接收者为值的也可以调用接收者为指针的
// 接收者为值的 只能调用接收者为值的

plst:=new(List)
Counter(plst,1,10)

fmt.Println(IsLarge(plst))

 course:=new(Course)
 course.Title="go lang in action"
 course.SubTitle="go 实战"

 fmt.Println(course)
 fmt.Printf("course= %v\n",course)


 data:=[]int{23,50,78,11,19,60}

 ia:=IntArray(data)

 Sort(ia)

 fmt.Println(ia)
}


func Sort(data sort.Interface){
	for pass:=1;pass<data.Len();pass++{
		for i:=0;i<data.Len()-pass;i++{
			if data.Less(i+1,i){
				data.Swap(i,i+1)
			}
		}
	}
}

type IntArray []int

func (ia IntArray) Len() int{
	return len(ia)
}

func(ia IntArray) Less(i,j int) bool{
	return ia[i]<ia[j]
}

func(ia IntArray) Swap(i,j int){
	ia[i],ia[j]=ia[j],ia[i]
}

