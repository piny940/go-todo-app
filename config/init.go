package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
	c = viper.New()
	filename := fmt.Sprintf("config/environments/%s.yaml", env)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	c.SetConfigType("yaml")

}

func GetConfig() *viper.Viper {
	return c
}
