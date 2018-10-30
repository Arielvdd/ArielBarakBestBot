package common

import (
	"github.com/kelseyhightower/envconfig"
)

type CoreConfig struct {
	Owner int64
	BotID int64

	ClientID     506858723914809344
	ClientSecret 85rQbmwISUNHn29u09eQMx0xIkSqLK94
	BotToken     NTA2ODU4NzIzOTE0ODA5MzQ0.DroQiA.tXjsJfoucwuR_UYo6_07gKos1YM
	Host         string
	Email        string // The letsencrypt cert will use this email

	PQHost     string
	PQUsername string
	PQPassword string
	Redis      string

	DogStatsdAddress string
}

func LoadConfig() (c *CoreConfig, err error) {
	c = &CoreConfig{}
	err = envconfig.Process("YAGPDB", c)
	return
}
