package windward

import (
	"fmt"
	"testing"
)

func TestB(t *testing.T) {
	type A struct {
		Age  int
		Name string
	}
	a := new(A)
	a.Name = "xx"
	fmt.Println(a)
}

func TestA(t *testing.T) {
	name := "./a.yml"
	w := GetWindward()
	w.InitConf([]string{name})
	//kind := w.GetKey(name, "kind")
	//fmt.Println("kind-->", kind)
	//fmt.Printf("w=%+v", w)

	SecA := w.GetKey(name, "SecA")
	fmt.Println("SecA-->", SecA) // SecA--> map[SecC:/A/B/c.go keyB:998]

	SecC := w.GetKey(name, "SecC")
	fmt.Println("SecC-->", SecC) // SecC--> /A/B/c.go

	keyA := w.GetKey(name, "keyA")
	fmt.Println("keyA-->", keyA) // keyA--> hello^
	//
	keyC := w.GetKey(name, "keyC")
	fmt.Println("keyC-->", keyC) // keyC--> 1.0.1~|888`888

	//SecD := w.GetKey(name, "SecD")
	//fmt.Println("SecD-->", SecD) // SecD--> xxx

	a := w.GetKey(name, "SecA.SecC")
	fmt.Println("SecA.SecC-->", a) // SecA.SecC--> /A/B/c.go

	b := w.GetKey(name, "SecA.keyB")
	fmt.Println("SecA.keyB-->", b) // SecA.keyB--> 998

	//c := w.GetKey(name, "SecA.e")
	//fmt.Println("----------")
	//fmt.Println("SecA.e-->", c) // SecA.keyB--> 998
}
