package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

//std | file
const (
	LogOutputStd  = "std"
	LogOutputFile = "file"
)

type config struct {
	HttpConf *httpConf `toml:"http"`
	DBConf   *dbConf   `toml:"db"`
	LogConf  *logConf  `toml:"log"`
}

type logConf struct {
	Level      string `toml:"level"`
	Output     string `toml:"output"`
	FilePath   string `toml:"filepath"`
	FileDirAbs string
	FileName   string
}

type httpConf struct {
	IP         string `toml:"ip"`
	Port       string `toml:"port"`
	Mode       string `json:"mode"`
	Endpoint   string `toml:"endpoint"`
	Passphrase string `toml:"passphrase"`
	RestServer string `toml:"restServer"`
	CAPath     string `toml:"capath"`
}

func (this *httpConf) Addr() string {
	return fmt.Sprintf("%s:%s", this.IP, this.Port)
}

type dbConf struct {
	IP       string `toml:"ip"`
	Port     string `toml:"port"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

func (this *dbConf) Uri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		this.UserName,
		this.Password,
		this.IP,
		this.Port,
		"admin",
	)
}

func (this *httpConf) GetEndpoint() string {
	return fmt.Sprintf("%s", this.Endpoint,
	)
}
func (this *httpConf) GetIp() string {
	return fmt.Sprintf("%s", this.IP,
	)
}
func (this *httpConf) GetCAPath() string {
	return fmt.Sprintf("%s", this.CAPath,
	)
}

func (this *httpConf) GetPassphrase() string {
	return fmt.Sprintf("%s", this.Passphrase,
	)
}

func (this *httpConf) GetRestServer() string {
	return fmt.Sprintf("%s", this.RestServer,
	)
}

var Config config

const configFile = "./config.toml"

//todo 相对路径

func init() {
	initFromFile(configFile)
}

func initFromFile(file string) {
	if _, err := toml.DecodeFile(file, &Config); err != nil {
		panic(err)
	}

	validateConfig()
	initLog()
}

func validateConfig() {
	if _, err := logrus.ParseLevel(Config.LogConf.Level); err != nil {
		panic(err)
	}

	if Config.LogConf.Output != LogOutputStd &&
		Config.LogConf.Output != LogOutputFile {
		panic("invalid log output")
	}

	if Config.LogConf.Output == LogOutputFile {
		if "" == strings.TrimSpace(Config.LogConf.FilePath) {
			panic("invalid log.filepath")
		}

		Config.LogConf.FileDirAbs, Config.LogConf.FileName = filepath.Split(Config.LogConf.FilePath)
	}

	if Config.DBConf.UserName == "" ||
		Config.DBConf.IP == "" ||
		Config.DBConf.Port == "" {

		panic("invalid db config")
	}

	if Config.HttpConf.Mode != gin.DebugMode &&
		Config.HttpConf.Mode != gin.ReleaseMode &&
		Config.HttpConf.Mode != gin.TestMode {

		panic("http mode invalid")
	}
}
