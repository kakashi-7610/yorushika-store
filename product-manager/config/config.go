package config

import (
	"fmt"
	"strings"
	"yorushika-store/product-manager/repository"
	"yorushika-store/product-manager/server"

	"github.com/spf13/viper"
)

type Configs struct {
	LogConfig    *LogConfig
	ServerConfig *server.ServerConfig
	DbConfig     *repository.DatabaseConfig
}

func NewConfigs() (*Configs, error) {
	v := viper.New()

	// 環境変数が設定されている場合は優先
	v.SetEnvPrefix("ypm")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// ref: https://qiita.com/takehanKosuke/items/1b17ade882b50cf2d737
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config/")
	err := v.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("failed read config file, error: %v", err)
		return nil, err
	}

	filename := v.GetString("log.filename")
	if len(filename) == 0 {
		err := fmt.Errorf("log filename is not configured")
		return nil, err
	}
	logConfig := NewLogConfig(filename)
	logConfig.LoggingSettings()

	srvPort := v.GetString("server.port")
	if len(srvPort) == 0 {
		err := fmt.Errorf("server port is not configured")
		return nil, err
	}
	serverConfig := server.NewServerConfig(srvPort)

	hostname := v.GetString("db.hostname")
	if len(hostname) == 0 {
		err := fmt.Errorf("db hostname is not configured")
		return nil, err
	}
	dbPort := v.GetString("db.port")
	if len(dbPort) == 0 {
		err := fmt.Errorf("db port is not configured")
		return nil, err
	}
	user := v.GetString("db.user")
	if len(user) == 0 {
		err := fmt.Errorf("db user is not configured")
		return nil, err
	}
	password := v.GetString("db.password")
	if len(password) == 0 {
		err := fmt.Errorf("db password is not configured")
		return nil, err
	}
	dbname := v.GetString("db.dbname")
	if len(dbname) == 0 {
		err := fmt.Errorf("db dbname is not configured")
		return nil, err
	}
	retry := v.GetInt("db.retry")
	if retry == 0 {
		err := fmt.Errorf("db retry is not configured")
		return nil, err
	}
	dbConfig := repository.NewDatabaseConfig(hostname, dbPort, user, password, dbname, retry)

	configs := &Configs{
		LogConfig:    logConfig,
		ServerConfig: serverConfig,
		DbConfig:     dbConfig,
	}

	return configs, nil
}
