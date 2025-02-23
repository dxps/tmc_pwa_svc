package config

type Config struct {
	Db struct {
		Driver string `env:"DB_DRIVER, default=postgres"`
		DSN    string `env:"DB_DSN, default=postgres://tmc:tmc@localhost:5457/tmc?sslmode=disable&timezone=Europe/Bucharest"`
	}
	Servers struct {
		FrontendPort int `env:"FRONTEND_PORT, default=8081"`
		BackendPort  int `env:"BACKEND_PORT, default=8082"`
	}
}
