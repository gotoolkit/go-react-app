package config

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Charset  string
}

func GetDefaultConfig() *DBConfig {
	return &DBConfig{
		Dialect:  "mysql",
		Username: "guest",
		Password: "password",
		Host:     "localhost",
		Port:     "3306",
		Name:     "todo",
		Charset:  "utf8",
	}
}
