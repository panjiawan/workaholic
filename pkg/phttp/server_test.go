package phttp

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestServer(t *testing.T) {
	s := &fasthttp.Server{
		MaxRequestBodySize: 1024,
	}
	server := fasthttprouter.New()
	server.GET("/test/v1", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("/test/v1 call")
	})
	s.Handler = server.Handler
	s.ListenAndServe(":8889")
}
