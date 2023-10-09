package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
	godotenv.Load(fmt.Sprintf(".env.%s", env))
	c = viper.New()
	yamlConfig := readYamlConfig(env)

	configStr := os.ExpandEnv(yamlConfig)

	c.SetConfigType("yaml")
	err := c.ReadConfig(strings.NewReader(configStr))
	if err != nil {
		panic(err)
	}
}

func GetConfig() *viper.Viper {
	return c
}

func readYamlConfig(env string) string {
	filename := fmt.Sprintf("config/environments/%s.yaml", env)
	configFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	fileContent := make([]byte, 1024)
	_, err = configFile.Read(fileContent)
	if err != nil {
		panic(err)
	}
	return string(fileContent)
}
