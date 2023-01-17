package global

import (
	"github.com/MultiMx/RancherRedeployAction/global/models"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var Config models.Config

func init() {
	if e := envconfig.Process("input", &Config); e != nil {
		log.Fatalln(e)
	}
}
