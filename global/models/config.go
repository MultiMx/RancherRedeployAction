package models

import "github.com/MultiMx/RancherRedeployAction/pkg/kube"

type Config struct {
	kube.Config
	WaitActive bool `envconfig:"wait"`
}
