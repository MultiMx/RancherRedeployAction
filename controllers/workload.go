package controllers

import (
	"fmt"
	"github.com/MultiMx/RancherRedeployAction/pkg/kube"
	log "github.com/sirupsen/logrus"
	"time"
)

func ReDeploy(api *kube.Kube) error {
	var e error
	var counter uint8
	for {
		if e = api.Redeploy(); e == nil {
			break
		}
		counter++
		if counter >= 5 {
			return e
		}
		log.Errorln("Request redeploy failed: ", e)
		time.Sleep(time.Second)
	}
	return nil
}

func WaitWorkloadAvailable(api *kube.Kube) error {
	var err = make(chan error)
	go func() {
		var counter uint8 = 0
		var ok bool
		var e error
		for {
			time.Sleep(time.Second)
			if ok, e = api.WorkloadActive(); e != nil {
				counter++
				if counter >= 5 {
					err <- e
					return
				}
				log.Warnf("Get workload status failed: %v", e)
				continue
			} else if ok {
				err <- nil
				return
			}
			counter = 0
		}
	}()
	select {
	case e := <-err:
		return e
	case <-time.After(time.Minute * 5):
		return fmt.Errorf("workload waiting timeout")
	}
}
