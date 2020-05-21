package main

import (
	"sort"
	"fmt"
)

func main(){


	//map 声明及初始化
	var m1 map[string] int
	m2:=make(map[string]int)
	m3:=map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}

	m2["d"]=12

	m1=map[string] int{
		"1":4,
		"2":5,
		"3":6,
	}

	fmt.Println(m1==nil)

	fmt.Println(m3)

	fmt.Println(m1)

	fmt.Println(m2)

	if val,ok:=m3["1"];ok{
		fmt.Printf("key in map and value %d\n",val)
	}else{
		fmt.Println("key not found")
	}

	for key:= range m1{
		fmt.Printf("key=%s,value=%d\n",key,m1[key])
	}

	for key,val:=range m1{
		fmt.Println(key,val)
	}

	delete(m1,"1")
	fmt.Println(m1)

	var keys [] string 

	for key:=range m1{
		keys=append(keys,key)
	}

	sort.Strings(keys)

	for _,k:=range keys{
		fmt.Println(m1[k])
	}
}