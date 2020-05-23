package main

import (
	"fmt"
	"io"
	"os"
)


type Speaker interface{
	Say(string)
}

func SpeakAlphabet(speaker Speaker){
	speaker.Say("abcdefghigk...xyz")
}

type Person struct{
	name string
}

type SpeakWriter struct{
	w io.Writer
}
 
func (sw *SpeakWriter) Say(message string) {
      io.WriteString(sw.w,message)   
}

func (p *Person) Say(message string){
	fmt.Println(p.name,":",message)

}
func main(){
   james:=new(Person)
   james.name="James"
   SpeakAlphabet(james)

   	

	console:=new(SpeakWriter)
	console.w=os.Stdout

	SpeakAlphabet(console)
}