package phttp

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/panjiawan/workaholic/pkg/plog"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"strings"
)

type Service struct {
	router        *fasthttprouter.Router
	server        *fasthttp.Server
	rateLimiter   *rate.Limiter
	https         bool
	httpsCertFile string
	httpsKeyFile  string
	host          string
	port          int
}

type Option func(opt *Service)

type Method string

func New(opts ...Option) *Service {
	r := &Service{
		router: fasthttprouter.New(),
	}
	r.server = &fasthttp.Server{}
	r.server.Handler = r.router.Handler
	for _, f := range opts {
		f(r)
	}

	return r
}

func (h *Service) Register(path, method string, f fasthttp.RequestHandler) {
	method = strings.ToLower(method)
	switch method {
	case MethodGet:
		h.router.GET(path, f)
	case MethodPost:
		h.router.POST(path, f)
	case MethodOptions:
		h.router.OPTIONS(path, f)
	case "head":
		h.router.HEAD(path, f)
	case "put":
		h.router.PUT(path, f)
	case "patch":
		h.router.PATCH(path, f)
	case "delete":
		h.router.DELETE(path, f)
	}

}

func (h *Service) RegisterWS(path string, msgType int, cb WSCallback) {
	ws := &WS{
		cb:      cb,
		msgType: msgType,
		pool:    make(map[int64]*WSClient),
	}
	h.router.GET(path, ws.wsHandle)
}

// Run 启动函数
func (h *Service) Run() error {

	var addr = ":8889"
	if h.port != 0 {
		addr = fmt.Sprintf("%s:%d", h.host, h.port)
	}
	plog.Info("start http server", zap.String("addr", addr))

	var err error = nil
	if h.https {
		err = h.server.ListenAndServeTLS(addr, h.httpsCertFile, h.httpsKeyFile)
	} else {
		err = h.server.ListenAndServe(addr)
	}

	if err != nil {
		plog.Error("start http server error", zap.Error(err))
	}

	return err
}
