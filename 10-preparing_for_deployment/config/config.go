package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	//
	var err error
	config = viper.New()

	//set file format
	config.SetConfigType("yaml")

	//set the config file name
	config.SetConfigName(env)

	//get app path
	path, _ := os.Executable()
	// get file path
	filePath := filepath.Dir(path)
	//config folder
	configFolder := fmt.Sprintf("%v/config/", filePath)

	// find file folder
	config.AddConfigPath(configFolder)

	//try to read file to see if error exist
	err = config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error file not found %w", err))
	}

}

func GetConfig(env string) *viper.Viper {
	Init(env)

	return config
}
