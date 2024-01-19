package config

import (
	"fmt"
	"github.com/alpha-omega-corp/services/types"
	"github.com/spf13/viper"
)

type Manager interface {
	Read(config string) (err error)
	Hosts() (c types.ConfigHosts, err error)

	UserService() (c types.ConfigUserService, err error)
	ConfigService() (c types.ConfigConfigService, err error)
	GithubService() (c types.ConfigGithubService, err error)
}

type Handler interface {
	Manager() Manager
}

type manager struct {
	handler *viper.Viper
}

type handler struct {
	Handler
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Manager() Manager {
	return newManager(viper.New())
}

func newManager(h *viper.Viper) Manager {
	return &manager{
		handler: h,
	}
}

func (m *manager) Read(key string) (err error) {
	err = m.handler.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", key)
	if err != nil {
		return
	}

	m.handler.SetConfigType("yaml")
	err = m.handler.ReadRemoteConfig()

	return
}

func (m *manager) Hosts() (c types.ConfigHosts, err error) {
	err = m.Read("config_hosts")
	err = m.handler.Unmarshal(&c)

	fmt.Print(err)
	return
}

func (m *manager) UserService() (c types.ConfigUserService, err error) {
	err = m.Read("config_user")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *manager) GithubService() (c types.ConfigGithubService, err error) {
	err = m.Read("config_github")
	err = m.handler.Unmarshal(&c)

	return
}

func (m *manager) ConfigService() (c types.ConfigConfigService, err error) {
	err = m.Read("config_config")
	err = m.handler.Unmarshal(&c)

	return
}
