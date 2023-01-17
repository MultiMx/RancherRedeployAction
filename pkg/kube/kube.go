package kube

import "strings"

type Kube struct {
	Conf *Config
}

func New(conf *Config) *Kube {
	if !strings.HasSuffix(conf.Backend, "/") {
		conf.Backend += "/"
	}
	return &Kube{
		Conf: conf,
	}
}
