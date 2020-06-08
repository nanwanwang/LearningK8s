package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

//发送http请求
func sendHttpRequestDemo(){

	req,err:=http.NewRequest(http.MethodGet,"http://www.baidu.com",nil)
	if err!=nil{
		panic(err)
	}

	req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	resp,err:=http.DefaultClient.Do(req)
	//resp,err:=http.Get("http://www.baidu.com")
	if err!=nil {
		panic(err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type User struct{
	 Name string `json:"name"`
	 Age int  `json:"age"`
}

func main() {
	//sendHttpRequestDemo();
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"Hello Golang\n")
	})

	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		user:=User{
			Name: "demon",
			Age: 23,
		}
		userJson,err:=json.Marshal(user)
		if err!=nil{
			http.Error(writer,err.Error(),http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-type","application/json")
		writer.Write(userJson)
	})

	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		image:=path.Join("images","golang.png")
		http.ServeFile(writer,request,image)
	})

	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		 user:=User{
		 	Name: "demon",
		 	Age: 23,
		 }
		 htmlFile:=path.Join("Templates","index.html")
		 tmpl,err:=template.ParseFiles(htmlFile)
		 if err!=nil{
		 	http.Error(writer,err.Error(),http.StatusInternalServerError)
		 	return
		 }
		 tmpl.Execute(writer,user)
	})
	//http server
	log.Fatal(http.ListenAndServe(":1234",nil))
}
