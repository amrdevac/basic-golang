package viperconfig

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootpath   = filepath.Join(filepath.Dir(b), "../../configs/")
)

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(rootpath)
	viper.SetConfigName("env.dev")
	errInit := viper.ReadInConfig()
	if errInit != nil {
		fmt.Println(errInit)
		return
	}
}

func GetConfig(value string) string {
	return viper.GetString(value)
}
