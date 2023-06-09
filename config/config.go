package config

type Config struct {
	Name string `json:"name"`
}

func GetConfig2() Config {
	return Config{
		Name: "MyApp",
	}
}
