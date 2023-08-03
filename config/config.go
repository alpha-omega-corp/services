package config

import (
	"embed"
	"gopkg.in/yaml.v3"
	"io/fs"
	"sync"
)

type DbConfig struct {
	ADDR string `yaml:"host"`
	NAME string `yaml:"name"`
	USER string `yaml:"user"`
	PASS string `yaml:"pass"`
}

type Config struct {
	HOST string   `yaml:"host"`
	DB   DbConfig `yaml:"db"`
}

func LoadConfig(fs fs.FS, env string) (*Config, error) {
	return readConfig(fs, env)
}

func MakeConfig(unwrapOnce *sync.Once, embed embed.FS, unwrapped fs.FS) fs.FS {
	unwrapOnce.Do(func() {
		fileSys, err := fs.Sub(embed, "envs")
		if err != nil {
			panic(err)
		}
		unwrapped = fileSys
	})

	return unwrapped
}

func readConfig(fileSys fs.FS, env string) (*Config, error) {
	b, err := fs.ReadFile(fileSys, env+".yaml")
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
