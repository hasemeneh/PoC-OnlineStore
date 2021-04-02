package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
)

type config struct {
	readfile readFileFunc
	fatalln  logger
}
type readFileFunc func(filename string) ([]byte, error)
type logger func(v ...interface{})

func New() *config {
	return &config{
		readfile: ioutil.ReadFile,
		fatalln:  log.Fatalln,
	}
}

func (c *config) Read() *models.MainConfig {
	cfg := models.MainConfig{}
	configByte, err := c.readfile("/etc/order-config/config.json")
	if err != nil {
		c.fatalln("Something Wrong Happen", err.Error())
	}
	err = json.Unmarshal(configByte, &cfg)
	if err != nil {
		c.fatalln("Something Wrong Happen", err.Error())
	}
	return &cfg
}
