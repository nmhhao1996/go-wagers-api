package config

import "github.com/caarlos0/env/v9"

// Config is the configuration for the application
type Config struct {
	HTTPServer HTTPServerConfig
	Logger     LoggerConfig
	MySQL      MySQLConfig
}

// HTTPServerConfig is the configuration for the HTTP server
type HTTPServerConfig struct {
	Port int    `env:"API_PORT" envDefault:"8080"`
	Mode string `env:"API_MODE" envDefault:"development"`
}

// LoggerConfig is the configuration for the logger
type LoggerConfig struct {
	Level    string `env:"LOGGER_LEVEL" envDefault:"debug"`
	Mode     string `env:"LOGGER_MODE" envDefault:"development"`
	Encoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

// MySQLConfig is the configuration for the MySQL database
type MySQLConfig struct {
	URI   string `env:"MYSQL_URI" envDefault:"root:mysql@tcp(waggers-db:3306)/go_wagers?parseTime=true&loc=Local"`
	Debug bool   `env:"MYSQL_DEBUG" envDefault:"true"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
