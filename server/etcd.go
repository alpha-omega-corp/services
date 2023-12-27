package server

import "github.com/spf13/viper"

func NewConfigurationManager(handler *viper.Viper, config string, s struct{}) (err error) {
	err = handler.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", config)
	if err != nil {
		return
	}

	viper.SetConfigType("yaml")

	err = viper.ReadRemoteConfig()
	err = viper.Unmarshal(&s)

	return
}
