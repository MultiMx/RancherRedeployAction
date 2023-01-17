package kube

import (
	"fmt"
)

type Config struct {
	Backend     string `required:"true"`
	Project     string `required:"true"`
	Namespace   string `required:"true"`
	Workload    string `required:"true"`
	BearerToken string `required:"true" envconfig:"token"`
}

func (a Config) DeploymentUrl() string {
	return a.Backend + fmt.Sprintf(
		"project/%s/workloads/deployment:%s:%s",
		a.Project,
		a.Namespace,
		a.Workload,
	)
}

type Request struct {
	Url   string
	Query map[string]interface{}
	Body  interface{}
}
