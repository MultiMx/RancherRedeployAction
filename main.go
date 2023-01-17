package main

import (
	"github.com/MultiMx/RancherRedeployAction/controllers"
	"github.com/MultiMx/RancherRedeployAction/global"
	"github.com/MultiMx/RancherRedeployAction/pkg/kube"
	log "github.com/sirupsen/logrus"
)

func main() {
	api := kube.New(&global.Config.Config)

	e := controllers.ReDeploy(api)
	if e != nil {
		log.Fatalln(e)
	}
	if global.Config.WaitActive {
		if e = controllers.WaitWorkloadAvailable(api); e != nil {
			log.Fatalln(e)
		}
	}

	log.Infoln("Success")
}
