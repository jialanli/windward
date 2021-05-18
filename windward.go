package windward

import (
	"encoding/json"
	lacia "github.com/jialanli/lacia/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

const (
	TypeList5 = "map[interface{}]interface{}"
	TypeList6 = "map[string]interface{}"
	TypeList7 = "[]interface{}"
	TypeList8 = "[]map[string]interface{}"
)

var sufList = []string{"yml", "yaml", "json"}
var typeSelect = []string{TypeList5, TypeList6, TypeList7, TypeList8}

type fileStat struct {
	status   bool
	path     string // ./path/conf.yml
	confName string // conf.yml
	confType string // yml
	conf     map[string]interface{}
	bs       []byte
}

type Wind struct {
	pathFamily []*fileStat
}

var windy *Wind

func GetWindward() *Wind {
	if windy == nil {
		return new(Wind)
	}

	return windy
}

func (w *Wind) InitConf(paths []string) {
	if len(paths) <= 0 {
		log.Panicf("config can't init, because len(paths) is 0")
		return
	}
	paths = checkRepeat(paths)
	var family []*fileStat
	for _, path := range paths {
		unit := strings.Split(path, "/")
		confName := unit[len(unit)-1]
		unitType := strings.Split(confName, ".")
		confType := unitType[len(unitType)-1]
		family = append(family, &fileStat{
			status:   false,
			path:     path,
			confName: confName,
			confType: confType,
			conf:     make(map[string]interface{}),
		})
	}

	w.pathFamily = family

	if err := w.setConf(); err != nil {
		log.Printf("config init err: %v", err.Error())
		return
	}
	//log.Printf("config set done: %+v", respectiveMap)
	log.Println("windward init success")
}

func checkRepeat(paths []string) (result []string) {
	result, _ = lacia.RepeatElementSidesString(paths)
	return
}

func (w *Wind) setConf() error {
	resetM()
	if respectiveMap == nil {
		respectiveMap = make(map[string]interface{})
	}
	for _, c := range w.pathFamily {
		if err := readConf(c); err != nil {
			log.Fatalf("readConf err:%v", err.Error())
			return err
		}
		c.status = true
		log.Printf("===> %s init done\n", c.confName)
		//confList = append(confList, c.conf)
		respectiveMap[c.path] = c.conf
	}

	checkIdenticalConf()
	return nil
}

func readConf(conf *fileStat) error {
	bs0, err := readFile(conf.path)
	if err != nil {
		return err
	}
	conf.bs = bs0
	bs := []byte(lacia.DeletePreAndSufSpace(string(bs0)))

	switch conf.confType {
	case "json":
		if err := json.Unmarshal(bs, &conf.conf); err != nil {
			return err
		}

	case "yml", "yaml":
		if err := yaml.Unmarshal(bs, &conf.conf); err != nil {
			return err
		}
	}

	return nil
}

func readFile(path string) (bs []byte, err error) {
	return ioutil.ReadFile(path)
}
