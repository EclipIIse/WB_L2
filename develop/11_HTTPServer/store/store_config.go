package store

type Config struct {
	DatabaseURL string
}

func NewDbConfig() *Config {
	return &Config{
		DatabaseURL: "/Users/koi/Desktop/Work/WBTech/WB_L2/develop/11_HTTPServer/one.json",
	}
}
