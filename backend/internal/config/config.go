package config

import (
	"flag"

	"github.com/joho/godotenv"
	// "log"
	// "github.com/joho/godotenv"
)

type Config struct {
	Port int
	env  string
	DB   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	// Limiter struct {
	// 	Enabled bool
	// 	RPS     float64
	// 	Burst   int
	// }
	// SMTP struct {
	// 	Host     string
	// 	Port     int
	// 	Username string
	// 	Password string
	// 	Sender   string
	// }
	// CORS struct {
	// 	TrustedOrigins []string
	// }
}

func Load() Config {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	flag.IntVar(&cfg.Port, "port", 3000, "server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable", "PostgreSQL DSN")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	return cfg
}
