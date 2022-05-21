package conf

import (
	"encoding/json"
	"io/ioutil"
)

type WebConfig struct {
	Addr string `json:"addr"`
}

type PgConfig struct {
	Dsn          string `json:"dsn"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type Config struct {
	Web *WebConfig `json:"web"`
	Pg  *PgConfig  `json:"pg"`
}

var Cfg *Config

func InitFrom(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	Cfg = new(Config)
	if err := json.Unmarshal(buf, Cfg); err != nil {
		return err
	}
	return nil
}

func Init() error {
	return InitFrom("./config.json")
}
