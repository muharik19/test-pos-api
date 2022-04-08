package common

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	RootURL    string `split_words:"true" default:"/api.pos-v1"`
	Port       string `split_words:"true" default:"3030"`
	DBHost     string `split_words:"true" default:""`
	DBPort     string `split_words:"true" default:""`
	DBUser     string `split_words:"true" default:""`
	DBPassword string `split_words:"true" default:""`
	DBName     string `split_words:"true" default:"pos_db"`
	JwtKey     string `split_words:"true" default:"M4h35a*)&@"`
	JwtExpired string `split_words:"true" default:"1d"`
}

var Config Configuration
var ctx = context.Background()

func InitConfig() {
	err := envconfig.Process("pos", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithField("config", Config).Info("Config successfully loaded")
}
