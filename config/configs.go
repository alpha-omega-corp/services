package config

type HostsConfig struct {
	Hosts []Host `mapstruct:"hosts"`
}

type Host struct {
	Host string `mapstruct:"host"`
	Dsn  string `mapstruct:"dsn"`
}

type GithubConfig struct {
	ClassicToken      string `mapstruct:"githubToken"`
	OrganizationUrl   string `mapstruct:"organizationUrl"`
	ContainerRegistry string `mapstruct:"containerRegistry"`
}

type AuthenticationConfig struct {
	JwtSecret string `mapstruct:"secret"`
}

type ServiceConfig struct {
	Dsn string `mapstruct:"dsn"`
}
