package global

import (
	"github.com/MultiMx/RancherRedeployAction/global/models"
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

var Config models.Config

func init() {
	if e := env.Parse(&Config, env.Options{
		Prefix: "INPUT_",
	}); e != nil {
		log.Fatalln(e)
	}
}
