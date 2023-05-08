package config_data

const (
	ConfigFolder = "profiles/"

	EnvironmentName    = "env"
	EnvironmentDefault = "dev"
	EnvFileNameDefault = ".env"
)

func SetEnvironmentName(name string) {
	if name == "" {
		return
	}

	_environmentName = name
}

func SetEnvironmentDefault(value string) {
	_environmentDefault = value
}

func SetEnvFileName(fileName string) {
	if fileName == "" {
		return
	}

	_envFileName = fileName
}

func GetEnvironmentName() string {
	return _environmentName
}

func GetEnvironmentDefault() string {
	return _environmentDefault
}

func GetEnvFileName() string {
	return _envFileName
}
