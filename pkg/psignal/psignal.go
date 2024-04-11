package psignal

import (
	"github.com/panjiawan/workaholic/pkg/plog"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Callback func(s os.Signal, args interface{})

type Handle struct {
	sigChan chan os.Signal
	set     map[os.Signal]Callback
}

func (h *Handle) Register(s os.Signal, f Callback) {
	if _, ok := h.set[s]; !ok {
		h.set[s] = f
		signal.Notify(h.sigChan, s)
	}
}

func (h *Handle) Listen() {
	plog.Info("isignal listen")
	for {
		if s, ok := <-h.sigChan; ok {
			if s == syscall.SIGURG {
				continue
			}
			plog.Info("catch signal", zap.String("signal", s.String()))
			if _, ok = h.set[s]; ok {
				plog.Info("call signal function", zap.String("signal", s.String()))
				h.set[s](s, nil)
			}
		} else {
			break
		}
	}
	plog.Info("signal exit")
}

func New() *Handle {
	return &Handle{
		sigChan: make(chan os.Signal, 1),
		set:     map[os.Signal]Callback{},
	}
}
