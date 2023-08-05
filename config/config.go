package config

import (
	"embed"
	"gopkg.in/yaml.v3"
	"io/fs"
	"sync"
)

var (
	//go:embed envs
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

type Config struct {
	HOST   string `yaml:"HOST"`
	AUTH   string `yaml:"AUTH"`
	DOCKER string `yaml:"DOCKER"`
}

func Get(env string) *Config {
	c, err := Load(Make(&unwrapFSOnce, embedFS, unwrappedFS), env)
	if err != nil {
		panic(err)
	}

	return c
}

func Load(fs fs.FS, env string) (*Config, error) {
	return Read(fs, env)
}

func Make(unwrapOnce *sync.Once, embed embed.FS, unwrapped fs.FS) fs.FS {
	unwrapOnce.Do(func() {
		fileSys, err := fs.Sub(embed, "envs")
		if err != nil {
			panic(err)
		}
		unwrapped = fileSys
	})

	return unwrapped
}

func Read(fileSys fs.FS, env string) (*Config, error) {
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
