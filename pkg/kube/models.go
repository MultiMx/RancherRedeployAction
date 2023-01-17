package kube

import (
	"fmt"
)

type Config struct {
	Backend     string
	Project     string
	Namespace   string
	Workload    string
	BearerToken string
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
