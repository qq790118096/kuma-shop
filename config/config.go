package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type _ConfigINF interface {
	Url() string
	Obj() interface{}
	SetTmpl(string)
	parse([]byte)
}

type Config struct {
	Tmpl           string
	UrlString      string
	Instance       interface{}
	Params         map[string]string
	ConfigFilename string
}

func (c *Config) Url() string {
	return c.UrlString
}
func (c *Config) Obj() interface{} {
	return c.Instance
}
func (c *Config) SetTmpl(tmpl string) {
	c.Tmpl = tmpl
}
func (c *Config) parse(jsonConfig []byte) {
	err := json.Unmarshal(jsonConfig, &c.Instance)
	c.Params = make(map[string]string)
	if err != nil {
		panic(err)
	}
	for k, v := range c.Instance.(map[string]interface{}) {
		c.Params[k] = v.(string)
	}
}

type MysqlConfig struct {
	Config
	Host string
	Port string
	DB   string
	User string
	PWD  string
}

func (mc *MysqlConfig) parse(jsonConfig []byte) {
	mc.Config.parse(jsonConfig)
	mc.Host = mc.Params["host"]
	mc.Port = mc.Params["port"]
	mc.DB = mc.Params["database"]
	mc.User = mc.Params["user"]
	mc.PWD = mc.Params["password"]
	mc.Tmpl = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true"
	mc.UrlString = fmt.Sprintf(mc.Tmpl, mc.User, mc.PWD, mc.Host, mc.Port, mc.DB)
}

func Init(config _ConfigINF, path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	config.parse(content)
}

func ConfigFactory(_type string) (_ConfigINF, error) {
	path, _ := os.Getwd()
	path += "/config/"
	var config _ConfigINF
	switch _type {
	case "mysql":
		{
			path += "mysql.config.json"
			config = &MysqlConfig{
				Config: Config{},
			}
			break
		}
	}
	Init(config, path)
	return config, nil
}
