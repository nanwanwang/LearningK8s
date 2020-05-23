package main

import (
	"fmt"
	"strconv"
)


type Book struct{
	id int
	title string
	author string
	subject string
}

func printBook(book *Book){
	fmt.Printf("id=%d,title=%s,author=%s,subject=%s\n",book.id,book.title,book.author,book.subject)
	book.id=1000
}

func (book *Book) String() string{
	return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s\n",book.id,book.title,book.author,book.subject)	
}

func (book *Book) GetTitle()string{
	return book.title
}

func NewBook (id int,title string,author string,subject string) *Book{
	return &Book{id,title,author,subject}
}


type Integer int

func (it Integer) String() string{
	return strconv.Itoa(int(it))
}

func main(){
 var b *Book
 b=new(Book)

  b.id=1
  b.author="demon"
  b.title="go lang"
  b.subject="learning"

  b1:=Book{
	  1002,
	  "hah",
	  "assd",
	  "c#",
  }


 
 fmt.Println(b1.String())

 fmt.Println(b1.GetTitle())


  fmt.Println(*b)



  printBook(&b1)
  fmt.Println(b1)

  fmt.Println("***********")

   b3:=NewBook(1,"11","22","33")
   fmt.Println(b3.String())


   fmt.Println("*********************")

   it:=Integer(100)

   str:=it.String()

   fmt.Printf("str=%s,type=%T\n",str,str)
}