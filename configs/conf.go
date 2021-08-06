package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	Conf = &config{}
)

func init() {
	confFile := flag.String("conf", "configs/dev.toml", "Please enter the profile address using -conf=xxx.toml.")
	flag.Parse()
	fmt.Println("using config : " + *confFile)
	if _, err := toml.DecodeFile(*confFile, Conf); err != nil {
		panic(err)
	}
}

type config struct {
	MySQL *MySQL `toml:"mysql"`
	App   *App   `toml:"app"`
	Zap   *Zap   `toml:"zap"`
}

type MySQL struct {
	Driver      string `toml:"driver"`
	Dsn         string `toml:"dsn"`
	MaxIdleConn int    `toml:"maxIdleConn"`
	MaxOpenConn int    `toml:"maxOpenConn"`
	MaxLifetime int    `toml:"maxLifetime"`
}

type App struct {
	Mode      string   `toml:"mode"`
	Name      string   `toml:"name"`
	Level     string   `toml:"level"`
	Port      int64    `toml:"port"`
	Secret    string   `toml:"secret"`
	UnAuthUrl []string `toml:"unAuthUrl"`
}

type Zap struct {
	File       string `toml:"file"`
	MaxSize    int    `toml:"maxSize"`
	MaxAge     int    `toml:"maxAge"`
	MaxBackups int    `toml:"maxBackups"`
	Compress   bool   `toml:"compress"`
	Level      string `toml:"level"`
}
