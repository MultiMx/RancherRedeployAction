package main

import (
	"github.com/MultiMx/RancherRedeployAction/controllers"
	"github.com/MultiMx/RancherRedeployAction/global"
	"github.com/MultiMx/RancherRedeployAction/pkg/kube"
	log "github.com/sirupsen/logrus"
)

func main() {
	api := kube.New(&global.Config.Config)

	err := controllers.ReDeploy(api)
	if err != nil {
		log.Fatalln(err)
	}
	if global.Config.WaitActive {
		if err = controllers.WaitWorkloadAvailable(api); err != nil {
			log.Fatalln(err)
		}
	}

	log.Infoln("Success")
}
