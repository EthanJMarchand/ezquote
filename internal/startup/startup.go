package startup

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	APIURL string
	Server string
}

func LoadEnv(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}
	config := &Config{
		APIKey: os.Getenv("APIKEY"),
		APIURL: os.Getenv("APIURL"),
		Server: os.Getenv("SERVER_ADDRESS"),
	}
	if config.APIKey == "" {
		return nil, errors.New("APIKey cannot be blank")
	}
	if config.APIURL == "" {
		return nil, errors.New("APIURL cannot be blank")
	}
	if config.Server == "" {
		return nil, errors.New("server address cannot be blank")
	}
	return config, nil
}
