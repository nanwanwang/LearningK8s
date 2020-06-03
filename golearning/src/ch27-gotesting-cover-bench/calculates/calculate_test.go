package calculates

import (
	"testing"
)

func  TestEven(t *testing.T) {
	//t.Log("Start TestEvent")
	//t.Error("This is a T.error function")
	if !Even(10){
		t.Log("10 must be even")
		t.Fail()
	}

	if !Even(78){
		t.Error("78 is not even")
	}
}


func TestTriangle(t *testing.T){
	tests:=[]struct{a,b,c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
		{30000,40000,50000},
	}

	for _,tt:=range tests{
		if actual:=calcTriangle(tt.a,tt.b);actual!=tt.c{
			t.Errorf("calcTriangle(%d,%d);"+
			"got %d,excepted %d",tt.a,tt.b,actual,tt.c)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	 input:=struct{
	 	a,b,c int
	 }{
	 	3001,4002,7003,
	 }
	 for i:=0;i<b.N;i++ {
		 if result := Add(input.a, input.b); result != input.c {
			 b.Errorf("calcTriangle(%d,%d);"+
				 "got %d,excepted %d", input.a, input.b, input.c, result)
		 }
	 }
}