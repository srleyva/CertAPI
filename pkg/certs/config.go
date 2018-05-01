package certs

import "github.com/spf13/viper"

type Config struct {
	Template string
	CA string
	PKI PKI
}

type PKI struct {
	Organization string
	OrganizationalUnit string
	Locality string
	Province string
	Country string
}

func NewConfig(filepathes []string) (*Config, error) {
	config := new(Config)
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	for _, path := range filepathes {
		v.AddConfigPath(path)
	}

	err := v.ReadInConfig()

	v.Unmarshal(&config)

	return config, err
}