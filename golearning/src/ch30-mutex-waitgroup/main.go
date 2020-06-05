package main

import (
	"fmt"
	"sync"
)

type SafeInt struct {
	a int
	lock sync.Mutex
}

func (si * SafeInt) Increase(){
	si.lock.Lock()
	defer si.lock.Unlock()
	si.a++
}

func (si *SafeInt) Get() int{
	si.lock.Lock()
	defer  si.lock.Unlock()
	return  si.a
}

func main() {

	 var wg sync.WaitGroup
	 var a SafeInt
	 a.Increase()
	for i:=0;i<5000;i++{
		wg.Add(1)
	 go func(){
	 	defer  wg.Done()
			a.Increase()
	 }()
	}
	wg.Wait()
	 //time.Sleep(time.Millisecond)
	 fmt.Println(a.Get())
}
