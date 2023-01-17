package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/RancherRedeployAction/modles"
	"github.com/MultiMx/RancherRedeployAction/util"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func request(Type string, c *modles.Request, query map[string]interface{}) (*http.Response, error) {
	res, e := util.Http.Request(Type, &tool.DoHttpReq{
		Url: c.Url,
		Header: map[string]interface{}{
			"User-Agent":    "curl/7.72.0",
			"Accept":        "*/*",
			"Authorization": "bearer " + c.BearerToken,
		},
		Query: query,
	})
	if e != nil {
		return nil, e
	}
	if res.StatusCode == 200 || res.StatusCode == 201 {
		return res, nil
	}
	d, e := io.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}
	return nil, fmt.Errorf("server throw error, http status %d : %s", res.StatusCode, string(d))
}

func redeployHandler(c *modles.Request) error {
	res, e := request("POST", c, map[string]interface{}{
		"action": "redeploy",
	})
	if e != nil {
		return e
	}
	_ = res.Body.Close()
	return nil
}

func workloadAvailable(c *modles.Request) (bool, error) {
	res, e := request("GET", c, nil)
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

func DoReDeploy() error {
	backend := os.Getenv("INPUT_BACKEND")
	if !strings.HasSuffix(backend, "/") {
		backend += "/"
	}
	config := &modles.Request{
		Url: backend + fmt.Sprintf(
			"project/%s/workloads/deployment:%s:%s",
			os.Getenv("INPUT_PROJECT"),
			os.Getenv("INPUT_NAMESPACE"),
			os.Getenv("INPUT_WORKLOAD"),
		),
		BearerToken: os.Getenv("INPUT_TOKEN"),
	}

	var e error
	var counter uint8
	for {
		if e = redeployHandler(config); e == nil {
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
				if ok, e = workloadAvailable(config); e != nil {
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
