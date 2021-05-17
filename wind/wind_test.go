package wind

import (
	"fmt"
	lacia "github.com/jialanli/lacia/utils"
	"testing"
)

func TestCLink(t *testing.T) {
	name := "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name})
	name0 := w.GetVal(name, "metadata.name")
	fmt.Println("metadata.name-->", name0) //

	namespace := w.GetVal(name, "metadata.namespace")
	fmt.Println("metadata.namespace-->", namespace) //

	ports := w.GetVal(name, "spec.ports")
	fmt.Println("spec.ports-->", ports) //

	port := w.GetVal(name, "spec.ports.port")
	fmt.Println("spec.ports.port-->", port) //

	protocol := w.GetVal(name, "spec.ports.protocol")
	fmt.Println("spec.ports.protocol-->", protocol) //

	targetPort := w.GetVal(name, "spec.ports.targetPort")
	fmt.Println("spec.ports.targetPort-->", targetPort) //

	sessionAffinity := w.GetVal(name, "spec.sessionAffinity")
	fmt.Println("spec.sessionAffinity-->", sessionAffinity) //

	types := w.GetVal(name, "spec.type")
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

	SecA := w.GetVal(name, "SecA")
	fmt.Println("SecA-->", SecA) // SecA--> map[SecC:/A/B/c.go keyB:998]

	SecC := w.GetVal(name, "SecC")
	fmt.Println("SecC-->", SecC) // SecC--> /A/B/c.go

	keyA := w.GetVal(name, "keyA")
	fmt.Println("keyA-->", keyA) // keyA--> hello^
	//
	keyC := w.GetVal(name, "keyC")
	fmt.Println("keyC-->", keyC) // keyC--> 1.0.1~|888`888

	SecD := w.GetVal(name, "SecD")
	fmt.Println("SecD-->", SecD) // SecD--> nihao

	a := w.GetVal(name, "SecA.SecC")
	fmt.Println("SecA.SecC-->", a) // SecA.SecC--> /A/B/c.go

	b := w.GetVal(name, "SecA.keyB")
	fmt.Println("SecA.keyB-->", b) // SecA.keyB--> 998

	c := w.GetVal(name, "SecB.keyA")
	fmt.Println("SecB.keyA-->", c) //
}

func TestMany(t *testing.T) {
	/*
		spec.ports.port--> 19987
		spec.ports.protocol--> TCP
		spec.ports.targetPort--> 19987
		spec.sessionAffinity--> None
		spec.type--> ClusterIP
		SecA--> map[SecC:/A/B/c.go keyB:998]
		SecC--> /A/B/c.go
		keyA--> hello^
		keyC--> 1.0.1~|888`888
		SecD--> nihao
		SecA.SecC--> /A/B/c.go
		SecA.keyB--> 998
		SecB.keyA--> hello^
	*/
	name, name1 := "./a.yml", "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name, name1})

	name0 := w.GetVal(name1, "metadata.name")
	fmt.Println("metadata.name-->", name0) //

	namespace := w.GetVal(name1, "metadata.namespace")
	fmt.Println("metadata.namespace-->", namespace) //

	ports := w.GetVal(name1, "spec.ports")
	fmt.Println("spec.ports-->", ports) //

	port := w.GetVal(name1, "spec.ports.port")
	fmt.Println("spec.ports.port-->", port) //

	protocol := w.GetVal(name1, "spec.ports.protocol")
	fmt.Println("spec.ports.protocol-->", protocol) //

	targetPort := w.GetVal(name1, "spec.ports.targetPort")
	fmt.Println("spec.ports.targetPort-->", targetPort) //

	sessionAffinity := w.GetVal(name1, "spec.sessionAffinity")
	fmt.Println("spec.sessionAffinity-->", sessionAffinity) //

	types := w.GetVal(name1, "spec.type")
	fmt.Println("spec.type-->", types) //

	SecA := w.GetVal(name, "SecA")
	fmt.Println("SecA-->", SecA) // SecA--> map[SecC:/A/B/c.go keyB:998]

	SecC := w.GetVal(name, "SecC")
	fmt.Println("SecC-->", SecC) // SecC--> /A/B/c.go

	keyA := w.GetVal(name, "keyA")
	fmt.Println("keyA-->", keyA) // keyA--> hello^
	//
	keyC := w.GetVal(name, "keyC")
	fmt.Println("keyC-->", keyC) // keyC--> 1.0.1~|888`888

	SecD := w.GetVal(name, "SecD")
	fmt.Println("SecD-->", SecD) // SecD--> nihao

	a := w.GetVal(name, "SecA.SecC")
	fmt.Println("SecA.SecC-->", a) // SecA.SecC--> /A/B/c.go

	b := w.GetVal(name, "SecA.keyB")
	fmt.Println("SecA.keyB-->", b) // SecA.keyB--> 998

	c := w.GetVal(name, "SecB.keyA")
	fmt.Println("SecB.keyA-->", c) //
}

func TestJs(t *testing.T) {
	name := "./s.json"
	w := GetWindward()
	w.InitConf([]string{name})

	c0 := w.GetVal(name, "name")
	fmt.Println("name-->", c0) //

	c1 := w.GetVal(name, "id")
	fmt.Println("id-->", c1) //

	c2 := w.GetVal(name, "type")
	fmt.Println("type-->", c2) //

	c3 := w.GetVal(name, "class")
	fmt.Println("class-->", c3) //

	c4 := w.GetVal(name, "class.data")
	fmt.Println("class.data-->", c4) //

	c5 := w.GetVal(name, "class.ins")
	fmt.Println("class.ins-->", c5) //

	c6 := w.GetVal(name, "other")
	fmt.Println("other-->", c6) //
}

func TestMany0(t *testing.T) {
	/*
		name--> 你好
		id--> 25
		type--> 哈哈
		class--> map[data:A ins:无名]
		class.data--> A
		class.ins--> 无名
		other--> [map[s1:雨落山岚 s2:55] map[s1:雨落山岚晚成风 s2:99]]
		metadata.name--> wordpress3
		metadata.namespace--> default
		spec.ports--> [map[port:19987 protocol:TCP targetPort:19987]]
		spec.ports.port--> 19987
		spec.ports.protocol--> TCP
		spec.ports.targetPort--> 19987
		spec.sessionAffinity--> None
		spec.type--> ClusterIP
	*/
	name, name1 := "./s.json", "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name, name1})

	c0 := w.GetVal(name, "name")
	fmt.Println("name-->", c0) //

	c1 := w.GetVal(name, "id")
	fmt.Println("id-->", c1) //

	c2 := w.GetVal(name, "type")
	fmt.Println("type-->", c2) //

	c3 := w.GetVal(name, "class")
	fmt.Println("class-->", c3) //

	c4 := w.GetVal(name, "class.data")
	fmt.Println("class.data-->", c4) //

	c5 := w.GetVal(name, "class.ins")
	fmt.Println("class.ins-->", c5) //

	c6 := w.GetVal(name, "other")
	fmt.Println("other-->", c6) //

	name0 := w.GetVal(name1, "metadata.name")
	fmt.Println("metadata.name-->", name0) //

	namespace := w.GetVal(name1, "metadata.namespace")
	fmt.Println("metadata.namespace-->", namespace) //

	ports := w.GetVal(name1, "spec.ports")
	fmt.Println("spec.ports-->", ports) //

	port := w.GetVal(name1, "spec.ports.port")
	fmt.Println("spec.ports.port-->", port) //

	protocol := w.GetVal(name1, "spec.ports.protocol")
	fmt.Println("spec.ports.protocol-->", protocol) //

	targetPort := w.GetVal(name1, "spec.ports.targetPort")
	fmt.Println("spec.ports.targetPort-->", targetPort) //

	sessionAffinity := w.GetVal(name1, "spec.sessionAffinity")
	fmt.Println("spec.sessionAffinity-->", sessionAffinity) //

	types := w.GetVal(name1, "spec.type")
	fmt.Println("spec.type-->", types) //
}

func TestUseManyConfigRead(t *testing.T) {
	/*
		JSON:err：<nil> 	 data：{Name:你好 ID:25 Type:哈哈 Class:{Data:A Ins:无名} Other:[{S1:雨落山岚 S2:55} {S1:雨落山岚晚成风 S2:99}]}
		YAML:err：<nil> 	date：{Kind:Service Metadata:{Name:wordpress3 Namespace:default} Spec:{Ports:[{TargetPort:19987 Port:19987 Protocol:TCP}] SessionAffinity:None Type:ClusterIP} ApiVersion:v1}--- PASS: TestUseManyConfigRead (0.01s)
	*/
	type other struct {
		S1 string `json:"s1"`
		S2 int    `json:"s2"`
	}
	type JsonDemo struct {
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Type  string `json:"type"`
		Class struct {
			Data string `json:"data"`
			Ins  string `json:"ins"`
		} `json:"class"`
		Other []other `json:"other"`
	}

	name, name1 := "./s.json", "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name, name1})

	//c6 := w.GetVal(name, "other")
	//fmt.Println("other-->", c6) //

	var data JsonDemo
	err := w.ReadConfig(name, &data)
	fmt.Printf("JSON:err：%v \t data：%+v", err, data)

	// yml
	type metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}

	type ports struct {
		TargetPort int    `yaml:"targetPort"`
		Port       int    `yaml:"port"`
		Protocol   string `yaml:"protocol"`
	}

	type spec struct {
		Ports           []ports `yaml:"ports"`
		SessionAffinity string  `yaml:"sessionAffinity"`
		Type            string  `yaml:"type"`
	}

	type YamlDemo struct {
		Kind       string   `yaml:"kind"`
		Metadata   metadata `yaml:"metadata"`
		Spec       spec     `yaml:"spec"`
		ApiVersion string   `yaml:"apiVersion"`
	}

	var dataYaml YamlDemo
	err = w.ReadConfig(name1, &dataYaml)
	fmt.Printf("\nYAML:err：%v \tdate：%+v", err, dataYaml)
}

func TestManyType(t *testing.T) {
	name, name1 := "./s.json", "./c.yml"
	w := GetWindward()
	w.InitConf([]string{name, name1})
	namespace := w.GetVal(name1, "metadata.namespace")
	fmt.Println("metadata.namespace-->", namespace) //

	port := w.GetVal(name1, "spec.ports.port")
	fmt.Println("spec.ports.port-->", port) //

	fmt.Println(lacia.GetValTypeOf(namespace))
	fmt.Println(lacia.GetValTypeOf(port))
}
