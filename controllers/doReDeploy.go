package controllers

import (
	"fmt"
	"github.com/MultiMx/RancherRedeployAction/pkg/kube"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func ReDeploy() error {
	api := kube.New(&kube.Config{
		Backend:     os.Getenv("INPUT_BACKEND"),
		Project:     os.Getenv("INPUT_PROJECT"),
		Namespace:   os.Getenv("INPUT_NAMESPACE"),
		Workload:    os.Getenv("INPUT_WORKLOAD"),
		BearerToken: os.Getenv("INPUT_TOKEN"),
	})

	var e error
	var counter uint8
	for {
		if e = api.Redeploy(); e == nil {
			break
		}

		log.Errorln("Request redeploy failed: ", e)

		counter++
		if counter >= 5 {
			return e
		}
		time.Sleep(time.Second)
	}

	if os.Getenv("INPUT_WAIT") == "true" {
		var err = make(chan error)
		go func() {
			counter = 0
			var ok bool
			for {
				time.Sleep(time.Second)
				if ok, e = api.WorkloadActive(); e != nil {
					log.Warnf("Get workload status failed: %v", e)
					counter++
					if counter == 5 {
						err <- e
						return
					}
					continue
				} else if ok {
					err <- nil
					return
				}
				counter = 0
			}
		}()
		select {
		case e = <-err:
			return e
		case <-time.After(time.Minute * 5):
			return fmt.Errorf("workload waiting timeout")
		}
	}

	return nil
}
