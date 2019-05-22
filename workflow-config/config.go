package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/mumushuiding/util"
)

// Configuration 数据库配置结构
type Configuration struct {
	Port           string
	ReadTimeout    int
	WriteTimeout   int
	DbLogMode      bool
	DbType         string
	DbName         string
	DbHost         string
	DbPort         string
	DbUser         string
	DbPassword     string
	DbMaxIdleConns int
	DbMaxOpenConns int
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
	err := Config.setFromEnv()
	if err != nil {
		panic(err)
	}
	// 打印配置信息
	config, _ := json.Marshal(&Config)
	log.Printf("configuration:%s", string(config))
}
func (c *Configuration) setFromEnv() error {
	// // 覆盖server配置
	// c.setServerFromEnv()
	// // 覆盖Db配置
	// c.setDbFromEnv()'
	// 获取对象Configuration的属性string流
	fieldStream, err := util.GetFieldChannelFromStruct(&Configuration{})
	if err != nil {
		return err
	}
	for fieldname := range fieldStream {
		if len(os.Getenv(fieldname)) > 0 {
			err = util.StructSetValByReflect(c, fieldname, os.Getenv(fieldname))
			if err != nil {
				return err
			}
		}
	}
	return nil
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
