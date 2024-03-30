package kube

import (
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/RancherRedeployAction/util"
	"io"
	"net/http"
)

func (a Kube) Request(Type string, req *Request) (*http.Response, error) {
	res, err := util.Http.Request(Type, &tool.DoHttpReq{
		Url: req.Url,
		Header: map[string]interface{}{
			"User-Agent":    "curl/7.72.0",
			"Accept":        "*/*",
			"Authorization": "bearer " + a.Conf.BearerToken,
		},
		Query: req.Query,
		Body:  req.Body,
	})
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 200 || res.StatusCode == 201 {
		return res, nil
	}
	d, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("server throw error, http status %d : %s", res.StatusCode, string(d))
}
