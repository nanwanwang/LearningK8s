package main

import "fmt"

func main(){
	var arr [3] int
	fmt.Println(arr)

	arr2:= [3] int {1,3,5}
	fmt.Println(arr2)

	arr3:=[...] int {1,3,6,6,5}
	fmt.Println(len(arr3))


	arr4:=[...]int{6,7,8,9,10}
	for i:=0;i<len(arr4);i++{
		fmt.Println(arr4[i])
	}

	arr5:=[...]int{1,2,3,4,5}
	for i,num := range arr5{
		fmt.Println(i,num)
	}

	for _,num :=range arr5{
		fmt.Println(num)
	}

	for i:=range arr4{
		fmt.Println(i)
	}


	arr6:=[2][3]int{
		{2,3,4},
		{5,6,7},
	}

	fmt.Println(arr6)


	arrTest(&arr4)
	fmt.Println(arr4)
}

func arrTest(arr *[5]int){
	(*arr)[0]=1000

}
