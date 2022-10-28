package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./")
	//viper.AddConfigPath("$HOME/")
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
}
