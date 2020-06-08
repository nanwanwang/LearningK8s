package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func timeDemo(){
	now:=time.Now()
	fmt.Println(now)
	//格式化时间只能用下面的时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	fmt.Println(now.Year(),now.Month(),now.Day())

	prev:=time.Date(2010,10,10,17,20,34,12340,time.UTC)

	fmt.Println(prev.Before(now))
	fmt.Println(prev.After(now))
	fmt.Println(prev.Equal(now))

	fmt.Println(now.Sub(prev))
}

func randDemo(){
	fmt.Println(rand.Intn(100))
}

func init(){
	rand.Seed(time.Now().Unix())
}

type ResponseData struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type BaseResponse struct{
	Code int `json:"name"`
	Data ResponseData `json:"data"`
}

func jsonDemo(){
	br:=BaseResponse{
		Code: 1,
		Data: ResponseData{
			Name: "demon",
			Age: 23,
		},
	}

	jsonBytes,_:=json.Marshal(&br)
	fmt.Println(string(jsonBytes))

	var br2 BaseResponse
	_=json.Unmarshal(jsonBytes,&br2)
	fmt.Println(br2)
}


func regexDemo(){
	input:="My email is 123@qq.com xxx@abc.com"
	exp,_:= regexp.Compile("([0-9a-zA-Z]+)@[0-9a-zA-Z]+.[0-9a-zA-Z]+")
	fmt.Println(exp.FindString(input))
	fmt.Println(exp.FindAllString(input,-1))

	for _,subMatch:= range exp.FindAllStringSubmatch(input,-1){
		fmt.Println(subMatch[1])
	}

    exp2:=	regexp.MustCompile("([0-9a-zA-Z]+)@[0-9a-zA-Z]+.[0-9a-zA-Z]+")
    fmt.Println(exp2.FindAllString(input,-1))
}
func main() {
	timeDemo()
	fmt.Println("================")
	randDemo()
	fmt.Println("============")
	jsonDemo()
	fmt.Println("================")
	regexDemo()
}
