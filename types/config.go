package types

type ConfigHosts struct {
	Config  ConfigHost `mapstruct:"config"`
	Gateway ConfigHost `mapstruct:"gateway"`
	Github  ConfigHost `mapstruct:"github"`
	User    ConfigHost `mapstruct:"user"`
}

type ConfigHost struct {
	Name string `mapstruct:"name"`
	Host string `mapstruct:"host"`
	Dsn  string `mapstruct:"dsn"`
}

type ConfigGithubOrg struct {
	Token    string `mapstruct:"token"`
	Name     string `mapstruct:"name"`
	Url      string `mapstruct:"url"`
	Registry string `mapstruct:"registry"`
}
