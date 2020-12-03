package config

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/crgimenes/goconfig"
	_ "github.com/crgimenes/goconfig/toml" // import required by goconfig
	"github.com/prest/prest/adapters"
	"github.com/prest/prest/adapters/mock"
	"github.com/prest/prest/adapters/postgres"
	pConf "github.com/prest/prest/config"
)

type Config struct {
	Cors      string           `toml:"cors" cfg:"cors"`
	Debug     bool             `toml:"debug" cfg:"debug" cfgDefault:"false"`
	PGHost    string           `toml:"pg_host" cfg:"pg_host" cfgDefault:"0.0.0.0"`
	PGPort    int              `toml:"pg_port" cfg:"pg_port" cfgDefault:"5432"`
	PGDBName  string           `toml:"pg_dbname" cfg:"pg_dbname" cfgDefault:"postgres"`
	PGUser    string           `toml:"pg_user" cfg:"pg_user" cfgDefault:"postgres"`
	PGPass    string           `toml:"pg_pass" cfg:"pg_pass" cfgDefault:"postgres"`
	Prest     PrestCore        `toml:"prest" cfg:"prest"`
	DBAdapter adapters.Adapter `toml:"-" cfg:"-"`
}

// PrestCore configuration
type PrestCore struct {
	Host       string `toml:"host" cfg:"host" cfgDefault:"127.0.0.1"`
	Port       int    `toml:"port" cfg:"port" cfgDefault:"8888"`
	Migrations string `toml:"migrations" cfg:"migrations" cfgDefault:"./migrations"`
}

var Get *Config

// Load is the loader of config envs
func Load() {
	if Get != nil {
		return
	}
	goconfig.File = "config.toml"
	Get = &Config{}
	err := goconfig.Parse(Get)
	if err != nil {
		log.Fatal("config.Load: ", err)
	}
	pConf.Load()
	cors := strings.Fields(Get.Cors)
	if len(cors) == 0 {
		cors = nil
	}
	if os.Getenv("PORT") == "" {
		pConf.PrestConf.HTTPPort = Get.Prest.Port
	} else {
		Get.Prest.Port = pConf.PrestConf.HTTPPort
	}
	if os.Getenv("DATABASE_URL") == "" {
		pConf.PrestConf.PGHost = Get.PGHost
		pConf.PrestConf.PGPort = Get.PGPort
		pConf.PrestConf.PGDatabase = Get.PGDBName
		pConf.PrestConf.PGUser = Get.PGUser
		pConf.PrestConf.PGPass = Get.PGPass
	} else {
		loadDefaultPGConfig()
	}
	pConf.PrestConf.Debug = Get.Debug
	pConf.PrestConf.CORSAllowOrigin = cors
	pConf.PrestConf.MigrationsPath = Get.Prest.Migrations
	pConf.PrestConf.EnableDefaultJWT = false
	pConf.PrestConf.PGMaxIdleConn = 0
	if os.Getenv("TEST") != "" {
		Get.DBAdapter = mock.New(&testing.T{})
	} else {
		postgres.Load()
		Get.DBAdapter = pConf.PrestConf.Adapter
	}
}

func loadDefaultPGConfig() {
	Get.PGHost = pConf.PrestConf.PGHost
	Get.PGPort = pConf.PrestConf.PGPort
	Get.PGDBName = pConf.PrestConf.PGDatabase
	Get.PGUser = pConf.PrestConf.PGUser
	Get.PGPass = pConf.PrestConf.PGPass
}
