package setting

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	RunMode string `toml:"RunMode"`
	Version string `toml:"Version"`
	App     struct {
		PAGESIZE  int    `toml:"PAGE_SIZE"`
		JWTSECRET string `toml:"JWT_SECRET"`
	} `toml:"app"`
	Server struct {
		HTTPPORT     int `toml:"HTTP_PORT"`
		READTIMEOUT  int `toml:"READ_TIMEOUT"`
		WRITETIMEOUT int `toml:"WRITE_TIMEOUT"`
	} `toml:"server"`
	Database struct {
		Type      string `toml:"Type"`
		Host      string `toml:"Host"`
		Username  string `toml:"Username"`
		Password  string `toml:"Password"`
		Dbname    string `toml:"Dbname"`
		Tablename string `toml:"Tablename"`
	} `toml:"database"`
}

var (
	cfg  *tomlConfig
	once sync.Once

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func Config() *tomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./ratblog/config/config.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

func init() {
	Config()
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = cfg.RunMode
}

func LoadServer() {
	sec := cfg.Server
	RunMode = cfg.RunMode
	HTTPPort = sec.HTTPPORT
	ReadTimeout = time.Duration(sec.READTIMEOUT) * time.Second
	WriteTimeout = time.Duration(sec.WRITETIMEOUT) * time.Second
}

func LoadApp() {
	sec := cfg.App
	JwtSecret = sec.JWTSECRET
	PageSize = sec.PAGESIZE
}
