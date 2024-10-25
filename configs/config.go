package configs

type Config struct {
	DB  *DBConfig
	App *AppConfig
}

func Configs() *Config {
	return &Config{
		DB:  LoadDBConfig(),
		App: LoadAppConfig(),
	}
}