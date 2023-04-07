package helper

type DBConfig struct {
	Host     string
	Username string
	Password string
	DB       string
}

type RedisConfig struct {
	Server   string
	Password string
}
