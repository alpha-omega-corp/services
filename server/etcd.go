package server

import (
	"github.com/alpha-omega-corp/services/config"
	"github.com/spf13/viper"
)

type ConfigManager interface {
}

type configManager struct {
	ConfigManager

	handler *viper.Viper
}

func NewConfigurationManager(handler *viper.Viper) ConfigManager {
	return &configManager{
		handler: handler,
	}
}

func (m *configManager) read(config string) (err error) {
	err = m.handler.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", config)
	if err != nil {
		return
	}

	m.handler.SetConfigType("yaml")
	err = m.handler.ReadRemoteConfig()

	return
}

func (m *configManager) HostConfig(svc string) (c config.HostConfig, err error) {
	err = m.read("/config/" + svc + ".yaml")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *configManager) AuthConfig() (c config.AuthenticationConfig, err error) {
	err = m.read("/config/auth.yaml")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *configManager) GithubConfig() (c config.GithubConfig, err error) {
	err = m.read("/config/github.yaml")
	err = m.handler.Unmarshal(&c)

	return
}
