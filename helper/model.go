package helper

type DBConfig struct {
	Host     string
	Username string
	Password string
	DB       string
	Port     int
}

type RedisConfig struct {
	Server   string
	Password string
	TTL      int
	Port     int
}
