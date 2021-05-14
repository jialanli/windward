package windward

import (
	"fmt"
	lacia "github.com/jialanli/lacia/utils"
	"log"
	"strings"
)

var confList []map[string]interface{}
var keyMap map[string]interface{}
var respectiveMap map[string]interface{}
var hasRepeatConf bool
var str []string

func resetM() {
	hasRepeatConf = false
	respectiveMap = nil
	keyMap = nil
	str = str[:0]
	confList = confList[:0]
}

func checkIdenticalConf() {
	if len(str) != 0 {
		str = str[:0]
	}
	if keyMap == nil {
		keyMap = make(map[string]interface{})
	}

	loopCheck(respectiveMap)

	log.Println("checkStr=", str)
	if res, _ := lacia.RepeatElementSidesString(str); len(res) < len(str) {
		hasRepeatConf = true
		log.Printf("\nhasRepeatConf is true, str=%v, res=%v", str, res)
	}
}

func loopCheck(m map[string]interface{}) {
	for k0, v := range m {
		if k0 == "ports" || k0 == "port" || k0 == "- port" {
			fmt.Printf("打印：k0=%v, v=%+v, type=%v\n", k0, v, lacia.GetValTypeOf(v))
		}
		//k := lacia.RemoveX(k0, "- ")
		k := k0
		if k0 == "port" || k0 == "- port" {
			log.Printf("key set before: k:%v, after:%v, v=%+v", k0, k, v)
		}

		keyMap[k] = v
		str = append(str, k)
		if str, ok := checkNeedScanByVertical(v); ok {
			switch str {
			case "map[interface{}]interface{}":
				innerV := v.(map[interface{}]interface{})
				loopCheck(convertM(innerV))
			case "map[string]interface{}":
				loopCheck(v.(map[string]interface{}))
			case "[]interface{}":
				loopCheck0(v.([]interface{}))
			default:
				return
			}
		}
	}
}

func loopCheck0(s []interface{}) {
	for i, v := range s {
		if v == "port" {
			fmt.Printf("打印：i=%v, v=%+v, type=%v\n", i, v, lacia.GetValTypeOf(v))
		}

		if str, ok := checkNeedScanByVertical(v); ok {
			switch str {
			case "map[interface{}]interface{}":
				innerV := v.(map[interface{}]interface{})
				loopCheck(convertM(innerV))
			case "map[string]interface{}":
				loopCheck(v.(map[string]interface{}))
			case "[]interface{}":
				loopCheck0(v.([]interface{}))
			default:
				return
			}
		}
	}
}

func convertM(v1 map[interface{}]interface{}) (v2 map[string]interface{}) {
	if v2 == nil {
		v2 = make(map[string]interface{})
	}
	for k, v := range v1 {
		kStr, ok := k.(string)
		if ok {
			v2[kStr] = v
		}
	}
	return
}

func checkNeedScanByVertical(m interface{}) (string, bool) {
	s := lacia.RemoveX(lacia.GetValTypeOf(m), " ")
	if s == "map[string]interface{}" {
		return s, true
	}
	if s == "[]map[string]interface{}" {
		return s, true
	}
	if s == "map[interface{}]interface{}" {
		return s, true
	}
	if s == "[]interface{}" {
		return s, true
	}

	return "", false
}

func setKeys() {

}

func getV(key string, v interface{}) (res interface{}) {
	if s, ok := checkNeedScanByVertical(v); ok {
		switch s {
		case "map[interface{}]interface{}":
			m := v.(map[interface{}]interface{})
			for k, next := range m {
				//fmt.Println("每一个：", k)
				//if k == "SecC" {
				//	fmt.Println("hhh")
				//}
				if key == k.(string) {
					res = next
					//fmt.Println("=============", res)
					return
				}

				res = getV(key, next)
				//fmt.Println("======dd=======", res)
			}
		case "map[string]interface{}":
			m := v.(map[string]interface{})
			for k, next := range m {
				if key == k {
					res = next
					return
				}

				res = getV(key, next)
				//fmt.Println("======dd0=======", res)
			}
		case "[]interface{}":
			inS := v.([]interface{})
			for _, val := range inS { // val:name: wordpress3
				getV(key, val)
			}
			//loopCheck0(v.([]interface{}))
		default:
			return
		}
	} else {
		return keyMap[key]
	}
	return
}

func getVLink(keys []string, v0 interface{}) (res interface{}) {
	for i, k := range keys {
		if _, ok := keyMap[k]; !ok {
			log.Fatalf("link: not found:'%s', please check", k)
			return
		}
		if s, ok := checkNeedScanByVertical(v0); ok {
			switch s {
			case "map[interface{}]interface{}":
				m := v0.(map[interface{}]interface{})
				for inK, next := range m {
					if inK.(string) != k {
						continue
					}

					if i == len(keys)-1 {
						return next
					}

					if i < len(keys)-1 {
						res = getVLink(keys[i:], next)
						return
					}
				}
			case "map[string]interface{}":
				m := v0.(map[string]interface{})
				for inK, next := range m {
					if inK != k {
						continue
					}

					if i == len(keys)-1 {
						return next
					}

					if i < len(keys)-1 {
						res = getVLink(keys[i:], next)
						return
					}
				}
			case "[]interface{}":
				inS := v0.([]interface{})
				for _, val := range inS { // val:name: wordpress3
					if i == len(keys)-1 {
						return val
					}

					if i < len(keys)-1 {
						res = getVLink(keys[i:], val)
						return
					}

					getVLink(keys[i+1:], val)
				}
			default:
				return
			}
		} else {
			return v0
		}
		//
		//if i != len(keys)-1 {
		//	res = getVLink(keys[i:], v0)
		//}
	}
	return
}

// find key by name
func (w *Wind) GetKey(name, key string) (res interface{}) {
	if strings.Contains(key, ".") {
		if v0, ok := respectiveMap[name]; ok {
			return getVLink(strings.Split(key, "."), v0)
		}
		return nil
	}

	for n, v := range respectiveMap {
		if n != name {
			continue
		}

		{
			if _, ok := keyMap[key]; !ok {
				log.Fatalf("not found:'%s', please check", key)
				return
			}

			res = getV(key, v)
			//fmt.Println("最后：", res)
		}
	}

	//if strings.Contains(key, ".") {
	//	keys = strings.Split(key, ".")
	//	for _, k := range keys {
	//		if _, ok := keyMap[k]; !ok {
	//			log.Fatalf("sub not found:'%s', please check", k)
	//			return
	//		}
	//	}
	//} else {
	//	if _, ok := keyMap[key]; !ok {
	//		log.Fatalf("not found:'%s', please check", key)
	//		return
	//	}
	//	for n, v := range respectiveMap {
	//		if n != name {
	//			continue
	//		}
	//
	//		res = getV(key, v)
	//	}
	//}

	return
}
