package config

import (
	"fmt"
	"strings"
	"yorushika-store/store-controller/server"

	"github.com/spf13/viper"
)

type Configs struct {
	LogConfig    *LogConfig
	ServerConfig *server.ServerConfig
}

func NewConfigs() (*Configs, error) {
	v := viper.New()

	// 環境変数が設定されている場合は優先
	v.SetEnvPrefix("ysc")
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

	port := v.GetString("server.port")
	if len(port) == 0 {
		err := fmt.Errorf("server port is not configured")
		return nil, err
	}
	static := v.GetString("server.static")
	if len(static) == 0 {
		err := fmt.Errorf("server static is not configured")
		return nil, err
	}
	serverConfig := server.NewServerConfig(port, static)

	configs := &Configs{
		LogConfig:    logConfig,
		ServerConfig: serverConfig,
	}

	return configs, nil
}
