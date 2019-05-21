package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// Configuration 数据库配置结构
type Configuration struct {
	Port           string `json:"SERVER_PORT"`
	ReadTimeout    int
	WriteTimeout   int64
	DbType         string `json:"DB_TYPE"`
	DbName         string `json:"DB_NAME"`
	DbHost         string `json:"DB_HOST"`
	DbPort         string `json:"DB_PORT"`
	DbUser         string `json:"DB_USER"`
	DbPassword     string `json:"DB_PASSWORD"`
	DbMaxIdleConns int    `json:"DB_MaxIdleConns"`
	DbMaxOpenConns int    `json:"DB_MaxOpenConns"`
}

// Config 数据库配置
var Config = &Configuration{}

func init() {
	LoadConfig()
}

// LoadConfig LoadConfig
func LoadConfig() {
	// 获取配置信息config
	Config.getConf()
	// 环境变量覆盖config
	Config.setFromEnv()
	// 打印配置信息
	config, _ := json.Marshal(&Config)
	log.Printf("configuration:%s", string(config))
}
func (c *Configuration) setFromEnv() *Configuration {
	// 覆盖server配置
	c.setServerFromEnv()
	// 覆盖Db配置
	c.setDbFromEnv()
	return c
}
func (c *Configuration) setServerFromEnv() *Configuration {
	if len(os.Getenv("SERVER_PORT")) > 0 {
		c.Port = os.Getenv("SERVER_PORT")
	}
	if len(os.Getenv("ReadTimeout")) > 0 {
		s, _ := strconv.Atoi(os.Getenv("ReadTimeout"))
		c.ReadTimeout = s
	}
	return c
}
func (c *Configuration) setDbFromEnv() *Configuration {
	if len(os.Getenv("DB_TYPE")) > 0 {
		c.DbType = os.Getenv("DB_TYPE")
	}
	if len(os.Getenv("DB_NAME")) > 0 {
		c.DbName = os.Getenv("DB_NAME")
	}
	if len(os.Getenv("DB_PORT")) > 0 {
		c.DbPort = os.Getenv("DB_PORT")
	}
	if len(os.Getenv("DB_HOST")) > 0 {
		c.DbHost = os.Getenv("DB_HOST")
	}
	if len(os.Getenv("DB_USER")) > 0 {
		c.DbUser = os.Getenv("DB_USER")
	}
	if len(os.Getenv("DB_PASSWORD")) > 0 {
		c.DbPassword = os.Getenv("DB_PASSWORD")
	}
	return c
}
func (c *Configuration) getConf() *Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("cannot open file config.json：%v", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Printf("decode config.json failed:%v", err)
		panic(err)
	}
	return c
}
