package config

type HostsConfig struct {
	Gateway   Host `mapstruct:"api_gateway"`
	AuthSVC   Host `mapstruct:"auth_svc"`
	DockerSVC Host `mapstruct:"docker_svc"`
}

type Host struct {
	Name string `mapstruct:"name"`
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
