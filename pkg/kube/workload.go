package kube

import (
	"encoding/json"
)

func (a Kube) Redeploy() error {
	res, err := a.Request("POST", &Request{
		Url: a.Conf.DeploymentUrl(),
		Query: map[string]interface{}{
			"action": "redeploy",
		},
	})
	if err != nil {
		return err
	}
	_ = res.Body.Close()
	return nil
}

func (a Kube) WorkloadActive() (bool, error) {
	res, err := a.Request("GET", &Request{
		Url: a.Conf.DeploymentUrl(),
	})
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	var resp struct {
		State string `json:"state"`
	}
	return resp.State == "active", json.NewDecoder(res.Body).Decode(&resp)
}
