package phttp

import (
	"github.com/fasthttp/websocket"
	"github.com/panjiawan/workaholic/pkg/plog"
	"go.uber.org/zap"
	"runtime/debug"
	"sync"
	"time"
)

type WSClient struct {
	sync.RWMutex
	ConnId   int64
	UserData interface{}
	conn     *websocket.Conn
	sendChan chan []byte
	msgType  int
	isClose  bool
}

func (w *WSClient) LoopWrite() {
	defer func() {
		if e := recover(); e != nil {
			plog.Info("loop write panic", zap.String("stack", string(debug.Stack())))
		}
	}()
	for msg := range w.sendChan {
		w.conn.WriteMessage(w.msgType, msg)
	}
}

func (w *WSClient) Send(msg []byte) {
	w.RLock()
	defer w.RUnlock()
	if w.isClose {
		return
	}
	select {
	case w.sendChan <- msg:
	case <-time.After(2 * time.Second):
		plog.Info("send chan time out")
	}
}
