package config

type HostsConfig struct {
	Gateway Host `mapstruct:"gateway"`
	Auth    Host `mapstruct:"auth"`
	Docker  Host `mapstruct:"docker"`
}

type Host struct {
	Name string `mapstruct:"name"`
	Host string `mapstruct:"host"`
	Dsn  string `mapstruct:"dsn"`
}

type GithubConfig struct {
	ClassicToken string       `mapstruct:"token"`
	Organization Organization `mapstruct:"organization"`
}

type Organization struct {
	Name     string `mapstruct:"name"`
	Url      string `mapstruct:"url"`
	Registry string `mapstruct:"registry"`
}

type AuthenticationConfig struct {
	JwtSecret string `mapstruct:"secret"`
}

type ServiceConfig struct {
	Dsn string `mapstruct:"dsn"`
}
