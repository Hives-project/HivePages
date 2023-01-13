package config

type KafkaConfig struct {
	Host     string `env:"KAFKA_HOST"`
	Port     string `env:"KAFKA_PORT"`
	User     string `env:"KAFKA_USER"`
	Password string `env:"KAFKA_PASSWORD"`
}

// Gets all values from the environment.
func (cfg *Config) loadKafkaConfig() KafkaConfig {
	envFields := cfg.loadEnvFields(KafkaConfig{})

	return KafkaConfig{
		Port:     envFields["Port"],
		Host:     envFields["Host"],
		User:     envFields["User"],
		Password: envFields["Password"],
	}
}
