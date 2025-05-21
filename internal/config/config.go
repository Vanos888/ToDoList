package config


type Config struct {
	PgDSN          string `envconfig:"PG_DSN"`
	RedisAddr      string `envconfig:"REDIS_ADDR"`
	RedisPass      string `envconfig:"REDIS_PASS"`
	RedisDB        int    `envconfig:"REDIS_DB"`
	JWTSecret      string `envconfig:"JWT_SECRET"`
	GrpcServerPort string `envconfig:"GRPC_SERVER_PORT"`
}

func (c *Config) GetJWTSecret() string {
	return c.JWTSecret
}
