package types

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	Database string
}

type JWTConfig struct {
	Secret             string
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Sender   string
}

type RedisConfig struct {
	Host string
}

type AppConfig struct {
	Port        int
	FrontendUrl string
	DatabaseConfig
	JWTConfig
	SMTPConfig
	RedisConfig
}
