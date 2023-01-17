package kube

import "encoding/json"

func (a Kube) Redeploy() error {
	res, e := a.Request("POST", &Request{
		Url: a.Conf.DeploymentUrl(),
		Query: map[string]interface{}{
			"action": "redeploy",
		},
	})
	if e != nil {
		return e
	}
	_ = res.Body.Close()
	return nil
}

func (a Kube) WorkloadActive() (bool, error) {
	res, e := a.Request("GET", &Request{
		Url: a.Conf.DeploymentUrl(),
	})
	if e != nil {
		return false, e
	}

	defer res.Body.Close()
	var resp struct {
		Selector struct {
			State string `json:"state"`
		} `json:"selector"`
	}
	return resp.Selector.State == "active", json.NewDecoder(res.Body).Decode(&resp)
}
