package phttp

import (
	"github.com/fasthttp/websocket"
	"github.com/panjiawan/workaholic/pkg/plog"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

type WSCallback interface {
	OnConnect(ctx *fasthttp.RequestCtx, client *WSClient) error
	OnMessage(cli *WSClient, msgType int, msg []byte)
	OnClose(client *WSClient)
}

type WS struct {
	sync.RWMutex
	cb      WSCallback
	incrId  int64
	msgType int
	pool    map[int64]*WSClient
}

var (
	MsgTypeText = websocket.TextMessage
	MsgTypeBin  = websocket.BinaryMessage
)

var (
	upGrader = websocket.FastHTTPUpgrader{
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			return true
		},
	}
)

func (w *WS) wsHandle(ctx *fasthttp.RequestCtx) {
	defer func() {
		if e := recover(); e != nil {
			plog.Error("[panic]", zap.String("stack", string(debug.Stack())))
		}
	}()

	var cli *WSClient = nil
	upGrader.Upgrade(ctx, func(c *websocket.Conn) {
		defer func() {
			if e := recover(); e != nil {
				plog.Error("[panic]", zap.String("stack", string(debug.Stack())))
			}
		}()
		cli = &WSClient{
			ConnId:   atomic.AddInt64(&w.incrId, 1),
			conn:     c,
			sendChan: make(chan []byte, 128),
			msgType:  w.msgType,
		}
		err := w.cb.OnConnect(ctx, cli)
		if err == nil {
			// start loop write
			go cli.LoopWrite()

			// add to pool
			w.Lock()
			w.pool[cli.ConnId] = cli
			w.Unlock()

			// read message
			for {
				msgType, message, err := c.ReadMessage()
				if err != nil {
					break
				}
				w.cb.OnMessage(cli, msgType, message)
			}
		}

		if cli != nil {
			// clear client data
			cli.Lock()
			cli.isClose = true
			close(cli.sendChan)
			defer func() {
				cli.Unlock()
				w.cb.OnClose(cli)
			}()

			// delete client from pool
			w.Lock()
			delete(w.pool, cli.ConnId)
			w.Unlock()
		}
	})
}
