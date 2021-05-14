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

	var str = "abc"
	s1 := fmt.Sprintf("1-%v", str)
	fmt.Println(s1)
	s2 := fmt.Sprintf("2-%v", str)
	fmt.Println(s2)

	biz := 1
	switch biz {
	case 2:
	case 3:
		//default:
		//	fmt.Println("1---")
	}
	fmt.Println("---")
}

func TestC(t *testing.T) {
	name := "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name})

	apiVersion := w.GetKey(name, "apiVersion")
	fmt.Println("apiVersion-->", apiVersion) //

	kind := w.GetKey(name, "kind")
	fmt.Println("kind-->", kind) //

	sessionAffinity := w.GetKey(name, "sessionAffinity")
	fmt.Println("sessionAffinity-->", sessionAffinity) //

	types := w.GetKey(name, "type")
	fmt.Println("type-->", types) //

	namespace := w.GetKey(name, "namespace")
	fmt.Println("namespace-->", namespace) //

	protocol := w.GetKey(name, "protocol")
	fmt.Println("protocol-->", protocol) //

	metadata := w.GetKey(name, "metadata")
	fmt.Println("metadata-->", metadata) //

	ports := w.GetKey(name, "ports")
	fmt.Println("ports-->", ports) //

	spec := w.GetKey(name, "spec")
	fmt.Println("spec-->", spec) //

	port := w.GetKey(name, "port")
	fmt.Println("port-->", port) //

	name0 := w.GetKey(name, "name")
	fmt.Println("name-->", name0) //
	/*
		apiVersion--> v1
		kind--> Service
		sessionAffinity--> None
		type--> ClusterIP
		namespace--> default
		protocol--> TCP
		metadata--> map[name:wordpress3 namespace:default]
		ports--> [map[port:19987 protocol:TCP targetPort:19987]]
		spec--> map[ports:[map[port:19987 protocol:TCP targetPort:19987]] sessionAffinity:None type:ClusterIP]
		port--> 19987
		name--> wordpress3
	*/
}

func TestCLink(t *testing.T) {
	name := "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name})
	name0 := w.GetKey(name, "metadata.name")
	fmt.Println("metadata.name-->", name0) //

	namespace := w.GetKey(name, "metadata.namespace")
	fmt.Println("metadata.namespace-->", namespace) //

	ports := w.GetKey(name, "spec.ports")
	fmt.Println("spec.ports-->", ports) //

	port := w.GetKey(name, "spec.ports.port")
	fmt.Println("spec.ports.port-->", port) //

	protocol := w.GetKey(name, "spec.ports.protocol")
	fmt.Println("spec.ports.protocol-->", protocol) //

	targetPort := w.GetKey(name, "spec.ports.targetPort")
	fmt.Println("spec.ports.targetPort-->", targetPort) //

	sessionAffinity := w.GetKey(name, "spec.sessionAffinity")
	fmt.Println("spec.sessionAffinity-->", sessionAffinity) //

	types := w.GetKey(name, "spec.type")
	fmt.Println("spec.type-->", types) //
	/*
	   metadata.name--> wordpress3
	   metadata.namespace--> default
	   spec.ports--> [map[port:19987 protocol:TCP targetPort:19987]]
	   spec.ports.port--> 19987
	   spec.ports.protocol--> TCP
	   spec.ports.targetPort--> 19987
	   spec.sessionAffinity--> None
	   spec.type--> ClusterIP
	*/
}

func TestA(t *testing.T) {
	name := "./a.yml"
	w := GetWindward()
	w.InitConf([]string{name})

	SecA := w.GetKey(name, "SecA")
	fmt.Println("SecA-->", SecA) // SecA--> map[SecC:/A/B/c.go keyB:998]

	SecC := w.GetKey(name, "SecC")
	fmt.Println("SecC-->", SecC) // SecC--> /A/B/c.go

	keyA := w.GetKey(name, "keyA")
	fmt.Println("keyA-->", keyA) // keyA--> hello^
	//
	keyC := w.GetKey(name, "keyC")
	fmt.Println("keyC-->", keyC) // keyC--> 1.0.1~|888`888

	SecD := w.GetKey(name, "SecD")
	fmt.Println("SecD-->", SecD) // SecD--> nihao

	a := w.GetKey(name, "SecA.SecC")
	fmt.Println("SecA.SecC-->", a) // SecA.SecC--> /A/B/c.go

	b := w.GetKey(name, "SecA.keyB")
	fmt.Println("SecA.keyB-->", b) // SecA.keyB--> 998

	c := w.GetKey(name, "SecB.keyA")
	fmt.Println("SecB.keyA-->", c) //

	//c := w.GetKey(name, "SecA.e")
	//fmt.Println("----------")
	//fmt.Println("SecA.e-->", c)
}
