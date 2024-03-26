package config

type DBEnv struct {
	Host     string `envconfig:"HOST"`
	Port     int    `envconfig:"PORT"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	DBName   string `envconfig:"DBNAME"`
}
