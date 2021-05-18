package windward

import (
	"encoding/json"
	"errors"
	"fmt"
	lacia "github.com/jialanli/lacia/utils"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
)

//var confList []map[string]interface{}
var keyMap map[string]interface{}
var respectiveMap map[string]interface{}
var hasRepeatConf bool
var str []string

func resetM() {
	hasRepeatConf = false
	respectiveMap = nil
	keyMap = nil
	str = str[:0]
	//confList = confList[:0]
}

func checkIdenticalConf() {
	if len(str) != 0 {
		str = str[:0]
	}
	if keyMap == nil {
		keyMap = make(map[string]interface{})
	}

	loopCheck(respectiveMap)

	if res, _ := lacia.RepeatElementSidesString(str); len(res) < len(str) {
		hasRepeatConf = true
		//log.Printf("\nhasRepeatConf is true, str=%v, res=%v", str, res)
	}
}

func loopCheck(m map[string]interface{}) {
	for k0, v := range m {
		k := k0
		//if k0 == "port" {
		//	log.Printf("key set before: k:%v, after:%v, v=%+v", k0, k, v)
		//}

		keyMap[k] = v
		str = append(str, k)
		if str, ok := checkNeedScanByVertical(v); ok {
			switch str {
			case TypeList5:
				innerV := v.(map[interface{}]interface{})
				loopCheck(convertM(innerV))
			case TypeList6:
				loopCheck(v.(map[string]interface{}))
			case TypeList7:
				loopCheck0(v.([]interface{}))
			default:
				return
			}
		}
	}
}

func loopCheck0(s []interface{}) {
	for _, v := range s {
		if str, ok := checkNeedScanByVertical(v); ok {
			switch str {
			case TypeList5:
				innerV := v.(map[interface{}]interface{})
				loopCheck(convertM(innerV))
			case TypeList6:
				loopCheck(v.(map[string]interface{}))
			case TypeList7:
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
	if lacia.ExistsInListString(typeSelect, s, true)[0] == -1 {
		return "", false
	}

	return s, true
}

func getV(key string, v interface{}) (res interface{}) {
	if s, ok := checkNeedScanByVertical(v); ok {
		switch s {
		case TypeList5:
			m := v.(map[interface{}]interface{})
			for k, next := range m {
				if key == k.(string) {
					res = next
					//fmt.Println("=============", res)
					return
				}

				res = getV(key, next)
				//fmt.Println("======dd=======", res)
			}
		case TypeList6:
			m := v.(map[string]interface{})
			for k, next := range m {
				if key == k {
					res = next
					return
				}

				res = getV(key, next)
				//fmt.Println("======dd0=======", res)
			}
		case TypeList7:
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
			case TypeList5:
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
			case TypeList6:
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
			case TypeList7:
				inS := v0.([]interface{})
				for _, val := range inS {
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
		//if i != len(keys)-1 {
		//	res = getVLink(keys[i:], v0)
		//}
	}
	return
}

func readConfig(path string, res interface{}, pathFamily []*fileStat) (err error) {
	var bs0 []byte
	var thisStat *fileStat
	for _, stat := range pathFamily {
		if stat.path != path {
			continue
		}
		bs0 = stat.bs
		thisStat = stat
	}

	if len(bs0) == 0 {
		return errors.New(fmt.Sprintf("no config file[%s] found, please check", path))
	}

	switch thisStat.confType {
	case "yml", "yaml":
		if err = yaml.Unmarshal(bs0, res); err != nil {
			return
		}
	case "json":
		if err = json.Unmarshal(bs0, res); err != nil {
			return
		}
	}

	return
}

// find key by name
func getVal(name, key string) (res interface{}) {
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

		if key == "" {
			return v
		}

		{
			if _, ok := keyMap[key]; !ok {
				log.Fatalf("not found:'%s', please check", key)
				return
			}

			res = getV(key, v)
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
