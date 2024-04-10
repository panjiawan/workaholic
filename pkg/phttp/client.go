package phttp

import (
	"errors"
	"github.com/valyala/fasthttp"
)

var (
	errStatus = errors.New("status error")
)

func Get(url string, data map[string]string, header map[string]string) ([]byte, error) {
	return Request(url, "GET", data, header)
}

func Post(url string, data map[string]string, header map[string]string) ([]byte, error) {
	return Request(url, "POST", data, header)
}

func Request(url, method string, param map[string]string, header map[string]string) ([]byte, error) {
	args := &fasthttp.Args{}
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseArgs(args)
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	for k, v := range param {
		args.Add(k, v)
	}
	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.SetBody(args.QueryString())

	client := &fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		return nil, err
	}

	if res.StatusCode() != fasthttp.StatusOK {
		return nil, errStatus
	}

	return res.Body(), nil
}
