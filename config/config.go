package config

import (
	"rtmp_stream_proxy/utils"
	"strings"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

const (
	ConfPath      = "conf"
	ConfName      = "app"
	ConfSuffix    = "yaml"
	ConfENVPrefix = "hsp"
)

type AppConf struct {
	App    App    `json:"app" yaml:"app" mapstructure:"app"`
	Server Server `json:"server" yaml:"server" mapstructure:"server"`
}

func init() {
	config = viper.New()
}

func InitConfig() (conf *AppConf, err error) {
	binDir := utils.GetBinDir()

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.SetEnvPrefix(ConfENVPrefix)
	config.AutomaticEnv()

	fullPath := utils.PathJoin(binDir, ConfPath, ConfName+"."+ConfSuffix)

	boo, _, _ := utils.Exist(fullPath)

	if boo {
		config.SetConfigName(ConfName)
		config.SetConfigType(ConfSuffix)
		config.AddConfigPath(utils.PathJoin(binDir, ConfPath))

		if err = config.ReadInConfig(); nil != err {
			return
		}
	}

	appConf := &AppConf{}
	if err = config.Unmarshal(appConf); nil != err {
		return
	}

	conf = appConf

	return
}
