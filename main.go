package main

import (
	"github.com/MultiMx/RancherRedeployAction/controllers"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	e := controllers.DoReDeploy()
	if e != nil {
		os.Exit(1)
	}
	logrus.Infoln("Success")
}
