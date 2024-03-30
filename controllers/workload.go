package controllers

import (
	"fmt"
	"github.com/MultiMx/RancherRedeployAction/pkg/kube"
	log "github.com/sirupsen/logrus"
	"time"
)

func ReDeploy(api *kube.Kube) error {
	var err error
	var counter uint8
	for {
		if err = api.Redeploy(); err == nil {
			break
		}
		counter++
		if counter >= 5 {
			return err
		}
		log.Errorln("Request redeploy failed: ", err)
		time.Sleep(time.Second)
	}
	return nil
}

func WaitWorkloadAvailable(api *kube.Kube) error {
	var errChan = make(chan error)
	go func() {
		var counter uint8 = 0
		var ok bool
		var err error
		for {
			time.Sleep(time.Second)
			if ok, err = api.WorkloadActive(); err != nil {
				counter++
				if counter >= 5 {
					errChan <- err
					return
				}
				log.Warnf("Get workload status failed: %v", err)
				continue
			} else if ok {
				errChan <- nil
				return
			}
			counter = 0
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(time.Minute * 5):
		return fmt.Errorf("workload waiting timeout")
	}
}
