package main

import (
	"reflect"
 	"fmt"
	// "ch17/book"
)





type TechBook struct{
	cat string
	int //匿名字段
	Book //匿名字段
}

type Book struct{ //标签
	Id int "书籍编号"
	Title string "书籍标题"
	Author string "书籍作者"
	Subject string "书籍主题"
}


func NewBook (id int,title string,author string,subject string) *Book{
	return &Book{id,title,author,subject}
}


func main(){
	book:=new(Book)

	book.Id=111
	book.Title="11"
	book.Author="11"
	book.Subject="55"


	book2:=NewBook(23,"12","34","45")
	fmt.Println(book)

	fmt.Println(*book2)
	RefTag(*book2,1)

	InitTechBook()


}

func (book *Book) String() string{
	return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s\n",book.Id,book.Title,book.Author,book.Subject)	
}

func InitTechBook(){

	   bk:=NewBook(1,"1","2","3")
	tb:=new(TechBook)
	tb.cat="tech"
	tb.int=2
	tb.Book=*bk

	fmt.Println("techBook cat=",tb.cat)
	fmt.Println("techBook int=",tb.int)
	fmt.Println("techBook book=",tb.Book.String())
	fmt.Println(tb.String())
}

func RefTag(b Book, idx int){
	bType:=reflect.TypeOf(b)
	inFiled:=bType.Field(idx)
	fmt.Printf("%v\n",inFiled.Tag)
}