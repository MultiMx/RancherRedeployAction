package controllers

import (
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/RancherRedeployAction/modles"
	"github.com/MultiMx/RancherRedeployAction/util"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

func handler(c *modles.Redeploy) error {
	Res, e := util.Http.PostRequest(&tool.DoHttpReq{
		Url: c.Url,
		Header: map[string]interface{}{
			"User-Agent":    "curl/7.72.0",
			"Accept":        "*/*",
			"Authorization": fmt.Sprintf("bearer %s:%s", c.AccessKey, c.SecretKey),
		},
		Query: map[string]interface{}{
			"action": "redeploy",
		},
	})
	if e != nil {
		return e
	}
	defer Res.Body.Close()
	if Res.StatusCode == 200 || Res.StatusCode == 201 {
		return nil
	}
	d, e := io.ReadAll(Res.Body)
	if e != nil {
		return e
	}
	return fmt.Errorf("server throw error, http status %d : %s", Res.StatusCode, string(d))
}

func DoReDeploy() error {
	backend := os.Getenv("INPUT_BACKEND")
	if !strings.HasSuffix(backend, "/") {
		backend += "/"
	}
	config := &modles.Redeploy{
		Url: backend + fmt.Sprintf(
			"project/%s/workloads/deployment:%s:%s",
			os.Getenv("INPUT_PROJECT"),
			os.Getenv("INPUT_NAMESPACE"),
			os.Getenv("INPUT_WORKLOAD"),
		),
		AccessKey: os.Getenv("INPUT_ACCESS_KEY"),
		SecretKey: os.Getenv("INPUT_SECRET_KEY"),
	}

	var e error
	var counter uint8
	for {
		if e = handler(config); e == nil {
			return nil
		}

		logrus.Errorln("Request redeploy failed: ", e)

		counter++
		if counter >= 5 {
			break
		}
		time.Sleep(time.Second)
	}
	return e
}
