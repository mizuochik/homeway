package environment

import (
	"errors"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	ServerURL *url.URL
}

func ReadFromEnv() (*Config, error) {
	var errms []string
	su := os.Getenv("SERVER_URL")
	if su == "" {
		errms = append(errms, "SERVER_URL is null")
	}
	suu, err := url.Parse(su)
	if err != nil {
		errms = append(errms, err.Error())
	}
	if errms != nil {
		return nil, errors.New(strings.Join(errms, ", "))
	}
	return &Config{
		ServerURL: suu,
	}, nil
}
