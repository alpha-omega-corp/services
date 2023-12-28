package server

import (
	"github.com/alpha-omega-corp/services/config"
	"github.com/spf13/viper"
)

type ConfigManager interface {
	Read(config string) (err error)
	HostsConfig() (c config.HostsConfig, err error)
	AuthConfig() (c config.AuthenticationConfig, err error)
	GithubConfig() (c config.GithubConfig, err error)
	ServiceConfig(svc string) (c config.ServiceConfig, err error)
}

type configManager struct {
	ConfigManager

	handler *viper.Viper
}

func NewConfigManager(handler *viper.Viper) ConfigManager {
	return &configManager{
		handler: handler,
	}
}

func (m *configManager) Read(config string) (err error) {
	err = m.handler.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", config)
	if err != nil {
		return
	}

	m.handler.SetConfigType("yaml")
	err = m.handler.ReadRemoteConfig()

	return
}

func (m *configManager) HostsConfig() (c config.HostsConfig, err error) {
	err = m.Read("/config/hosts.yaml")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *configManager) AuthConfig() (c config.AuthenticationConfig, err error) {
	err = m.Read("/config/auth.yaml")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *configManager) GithubConfig() (c config.GithubConfig, err error) {
	err = m.Read("/config/github.yaml")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *configManager) ServiceConfig(svc string) (c config.ServiceConfig, err error) {
	err = m.Read("/config/" + svc + ".yaml")
	err = m.handler.Unmarshal(&c)

	return
}
