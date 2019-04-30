package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Listen	string
	DBServers	map[string]DBServer
}

func UnmarshalConfig(tomlfile string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		return c, err
	}
	return c, nil
}

func (c Config) DBServerConf(key string) (DBServer, bool) {
	s, ok := c.DBServers[key]
	return s, ok
}

func (c Config) GetListenAddr() string {
	return c.Listen
}

type DBServer struct {
	Host string `toml:"host"`
	Port int `toml:"port"`
	DBName string `toml:"dbname"`
	User string `toml:"user"`
	Password string `toml:"password"`
}

func (m DBServer) ConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		m.User, m.Password, m.Host, m.Port, m.DBName)
	//return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	m.Host, m.Port, m.User, m.Password, m.DBName)
}
