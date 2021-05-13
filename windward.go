package windward

import (
	"encoding/json"
	lacia "github.com/jialanli/lacia/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

var sufList = []string{"yml", "yaml", "json"}

type fileStat struct {
	status   bool
	path     string // ./path/conf.yml
	confName string // conf.yml
	confType string // yml
	conf     map[string]interface{}
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
		log.Panicf("config can't init, because paths is 0")
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
	log.Printf("config set done: %+v", respectiveMap)
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
			log.Printf("readConf err:%v", err.Error())
			return err
		}
		c.status = true
		log.Printf("===> %s init done", c.confName)
		confList = append(confList, c.conf)
		respectiveMap[c.path] = c.conf
	}

	checkIdenticalConf()
	return nil
	// set configMap
	//if configMap == nil {
	//	configMap = map[string]interface{}{}
	//}
	//for i := range confList {
	//	for k, v := range confList[i] {
	//		configMap[k] = v
	//	}
	//}
}

func readConf(conf *fileStat) error {
	bs0, err := readFile(conf.path)
	if err != nil {
		return err
	}

	//bs0 = bytes.TrimPrefix(bs0, []byte{239, 187, 191}) //[]byte("\xef\xbb\xbf") Or []byte{239, 187, 191}
	log.Printf("type=%v", conf.confType)
	bs1 := lacia.DeletePreAndSufSpace(string(bs0))
	bs := []byte(bs1)
	//reader := bytes.NewReader(bs0)
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(reader)
	//bs := buf.Bytes()
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
