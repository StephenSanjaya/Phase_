package config

type DBEnv struct {
	DBName     string `envconfig:"NAME"`
	DBPort     int    `envconfig:"PORT"`
	DBHost     string `envconfig:"HOST"`
	DBUsername string `envconfig:"USERNAME"`
	DBPassword string `envconfig:"PASSWORD"`
}
