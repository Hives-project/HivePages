package config

type GRPCConfig struct {
	Host string `env:"HTTP_HOST"`
	Port string `env:"HTTP_PORT"`
}

// Gets all values from the environment.
func (cfg *Config) LoadGRPCConfig() GRPCConfig {
	envFields := cfg.loadEnvFields(GRPCConfig{})

	return GRPCConfig{
		Host: envFields["Host"],
		Port: envFields["Port"],
	}
}
