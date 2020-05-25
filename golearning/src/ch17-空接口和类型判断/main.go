package main

import "fmt"

type Phone interface{
	Call(string)
}

type Camera interface{
	Take() string
}

type SmartPhone interface{
	Phone
	Camera
	Download(string) string
}

func ListSmartPhoneFunction(sp SmartPhone){

 if v,ok:=sp.(*MiPhone);ok{
	 v.Call("10010")
	 fmt.Println("sp.Take()",v.Take())
	 fmt.Println("sp.Download()",v.Download("jjjjj"))
  }else{
	  fmt.Println("not Mi Phone pointer type")
  }
}

type MiPhone struct{
	Logo string
}


func (mp *MiPhone) Call(phone string){
	fmt.Println("call to phone:",phone)
}

func (mp *MiPhone) Take() string{
	return "logo.png"
}

func (mp *MiPhone) Download(url string) string{
    return fmt.Sprintf("visit url: %s",url)
}

//空接口 ,可以用interface{} 表示任何的数据类型
type Any interface{}


//type-switch
func GetAnyType(any interface{}){
	switch t:=any.(type){
	case int:
		fmt.Println("any is int type")
	case string:
		fmt.Println("any is string type")
	case *MiPhone:
		fmt.Println("any is MiPhone pointer type")
	default:
		fmt.Println("Unexpected type %T\n",t)
	}
}

type Queue [] interface{} 


func (queue *Queue) Push(n interface{} ){
	*queue=append(*queue,n)
}

func (queue *Queue) Pop() interface{} {
	head:=(*queue)[0]
	*queue=(*queue)[1:]
	return head
}

func main(){
  var sp SmartPhone
  mp:=new(MiPhone)
  mp.Logo="mi"
  sp=mp
  ListSmartPhoneFunction(sp)

  var  va1 Any
  va1=5
  fmt.Printf("va1 value:%v\n",va1)
  GetAnyType(va1)

  str:="111"
  va1=str
  fmt.Printf("va1 value:%v\n",va1)
  GetAnyType(va1)

  va1=*mp
  fmt.Printf("va1 vaule:%v\n",va1)
  GetAnyType(va1)

  queue:=Queue{1,2,3,4}
  queue.Push("s")
  fmt.Println(queue)

  queue.Pop()
  fmt.Println(queue)
}