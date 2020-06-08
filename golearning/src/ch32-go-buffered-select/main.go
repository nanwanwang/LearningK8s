package main

import (
	"fmt"
	"time"
)

func buffeedChannelDemo() {
	c := make(chan int, 3)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	c <- 10
	c <- 2
	c <- 5
	c <- 6
}

func selectDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	timeout := make(chan bool)
	go func() {
		time.Sleep(time.Millisecond)
		timeout <- true
	}()

	go func() {
		ch1 <- 10
		ch2<-20
	}()

	for {
		select {
		case <-ch1:
			fmt.Println("receive data from ch1")
		case <-ch2:
			fmt.Println("receive data from ch2")
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("no data")
		}
	}
}

//验证通道是否满了 使用select语句和default case来实现非阻塞.
func checkChannelFull(c chan int){
	select{
	case c<- 2:
		default:
			fmt.Println("channel is full")
	}
}

//init 函数是最先执行的
func  init()  {
	fmt.Println("init")
}
func main() {
	//buffeedChannelDemo()
	//time.Sleep(time.Millisecond)

	//selectDemo()
	c:=make(chan int,2)
	c<-1
	c<-2
	checkChannelFull(c)

	<-time.After(time.Second)
}
