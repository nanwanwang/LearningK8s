package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10] int
	for i:=0;i<10;i++{
		go func(i int) {
			for{
				a[i]++
				runtime.Gosched()
			}

			//fmt.Printf("I am from goroutine %d\n",i)
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
