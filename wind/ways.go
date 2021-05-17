package wind

import (
	"github.com/spf13/cast"
	"time"
)

func (w *Wind) GetVal(name, key string) interface{} {
	return getVal(name, key)
}

func (w *Wind) GetValString(name, key string) string {
	return cast.ToString(getVal(name, key))
}

func (w *Wind) GetValInt(name, key string) int {
	return cast.ToInt(getVal(name, key))
}

func (w *Wind) GetValUInt(name, key string) uint {
	return cast.ToUint(getVal(name, key))
}

func (w *Wind) GetValBool(name, key string) bool {
	return cast.ToBool(getVal(name, key))
}

func (w *Wind) GetValFloat64(name, key string) float64 {
	return cast.ToFloat64(getVal(name, key))
}

func (w *Wind) GetValTime(name, key string) time.Time {
	return cast.ToTime(getVal(name, key))
}

func (w *Wind) ReadConfig(path string, res interface{}) (err error) {
	if err = readConfig(path, res, w.pathFamily); err != nil {
		return
	}
	return
}
