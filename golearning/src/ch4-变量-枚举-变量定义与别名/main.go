package main


import(
	"fmt"
	"math"
)

//可导出的MAX
const(
	s1="golang"
	MAX=10
)
func constDemo(){
	const s string = "Hello"
	const a ,b=3,4
	fmt.Println(s,a,b)
	var c int 
	c= int(math.Sqrt(a*a+b*b))
	fmt.Println(c)
	fmt.Println(s1,MAX)

	
}

func enumDemo(){
	const(
		Monday=1+iota
		Tuesday
		Wednesday
		Tursday
		Friday
		Staurday
		Sunday
	)
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Tursday,Friday,Staurday)
}

func typeDefAndAlias(){
	type MyInt1 int //基于int定义的一个新类型,用的比较多
	type MyInt2 =int //MyInt2 和int 就是完全一样的 type alias
	var i int =1
	var i1 MyInt1=MyInt1(i)
	var i2 MyInt2=i
	fmt.Println(i1,i2)
}

func main(){

	constDemo()

	enumDemo()

	typeDefAndAlias()
}