package config

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
