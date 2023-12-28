package config

type GithubConfig struct {
	ClassicToken      string `mapstruct:"githubToken"`
	OrganizationUrl   string `mapstruct:"organizationUrl"`
	ContainerRegistry string `mapstruct:"containerRegistry"`
}

type AuthenticationConfig struct {
	JwtSecret string `mapstruct:"secret"`
}

type HostConfig struct {
	Host string `mapstruct:"host"`
	Dsn  string `mapstruct:"dsn"`
}
