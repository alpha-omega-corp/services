package types

type ConfigGithubService struct {
	Organization ConfigGithubOrg `mapstruct:"organization"`
}
type ConfigUserService struct {
	JwtSecret string `mapstruct:"secret"`
}

type ConfigConfigService struct {
	Mode string `mapstruct:"mode"`
}

type Environment struct {
	ServiceVar    string `mapstruct:"service"`
	RepositoryVar string `mapstruct:"repository"`
	ConfigVar     string `mapstruct:"config"`
}
