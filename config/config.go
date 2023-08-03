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

type DbConfig struct {
	ADDR string `yaml:"HOST"`
	NAME string `yaml:"NAME"`
	USER string `yaml:"USER"`
	PASS string `yaml:"PASS"`
}

type Config struct {
	HOST   string `yaml:"HOST"`
	AUTH   string `yaml:"AUTH"`
	DOCKER string `yaml:"DOCKER"`

	DB DbConfig `yaml:"DB"`
}

func Get(env string) (*Config, error) {
	return load(config(&unwrapFSOnce, embedFS, unwrappedFS), env)
}

func load(fs fs.FS, env string) (*Config, error) {
	return read(fs, env)
}

func config(unwrapOnce *sync.Once, embed embed.FS, unwrapped fs.FS) fs.FS {
	unwrapOnce.Do(func() {
		fileSys, err := fs.Sub(embed, "envs")
		if err != nil {
			panic(err)
		}
		unwrapped = fileSys
	})

	return unwrapped
}

func read(fileSys fs.FS, env string) (*Config, error) {
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
