package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig(configFileName string) *viper.Viper {
	viperInstance := viper.New()
	if configFileName == "" {
		configFileName = CONFIG_FILE_NAME
	}
	//设置配置文件
	viperInstance.SetConfigName(configFileName)
	viperInstance.SetConfigType(CONFIG_FILE_TYPE)
	//设置配置文件查找目录
	viperInstance.AddConfigPath(HOME_CONFIG_PATH)
	viperInstance.AddConfigPath(OPT_CONFIG_PATH)
	viperInstance.AddConfigPath(EMBED_CONFIG_PATH)
	//读取配置文件
	if err := viperInstance.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("file not found %v\n", err)
		} else {
			fmt.Printf("read cpnfig error %v\n", err)
		}
	}
	//监听配置文件变动
	viperInstance.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viperInstance.WatchConfig()

	/*v := viperInstance.Get("port")
	fmt.Println(v)

	if homePath, err := homedir.Dir(); err == nil {
		fmt.Printf("homedir is %s\n", homePath)
	}*/

	return viperInstance
}
