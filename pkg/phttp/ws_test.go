package phttp

import (
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/panjiawan/workaholic/pkg/plog"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

type WsCallback struct {
}

var cb = &WsCallback{}

func TestWs(t *testing.T) {
	plog.Start("./tmp/", "ws_test.log", true, true)

	handle := New(
		WithAddress("", 8899),
	)

	handle.RegisterWS("/", websocket.TextMessage, cb)

	if err := handle.Run(); err != nil {
		t.Error(err)
	}
}

func (w *WsCallback) OnConnect(ctx *fasthttp.RequestCtx, client *WSClient) error {
	// login
	fmt.Println("OnConnect")
	return nil
}

func (w *WsCallback) OnMessage(client *WSClient, msgType int, msg []byte) {
	fmt.Println("OnMessage")
	time.Sleep(3 * time.Second)
	panic("111")
}

func (w *WsCallback) OnClose(client *WSClient) {
	fmt.Println("OnClose")
}
