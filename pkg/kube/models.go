package kube

import (
	"fmt"
)

type Config struct {
	Backend     string `env:"BACKEND,required"`
	Project     string `env:"PROJECT,required"`
	Namespace   string `env:"NAMESPACE,required"`
	Workload    string `env:"WORKLOAD,required"`
	BearerToken string `env:"TOKEN,required"`
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
