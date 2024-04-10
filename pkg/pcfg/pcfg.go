package pcfg

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/panjiawan/workaholic/pkg/plog"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

type CfgType int

type service struct {
	sync.RWMutex
	cache map[string]interface{}
}

var (
	CfgTypeYaml CfgType = 1 //yaml
	CfgTypeJson CfgType = 2 //json
)

var handle *service

func init() {
	handle = &service{
		cache: map[string]interface{}{},
	}
}

func Load(t CfgType, key, path string, data interface{}) error {
	cfg, err := os.ReadFile(path)
	if err != nil {
		plog.Error("get cfg file error", zap.String("path", path), zap.Error(err))
		return err
	}
	err = errors.New("register failure")
	if t == CfgTypeYaml {
		err = yaml.Unmarshal(cfg, data)
		if err != nil {
			plog.Error("yaml Unmarshal file error", zap.String("path", path), zap.Error(err))
		}
	} else if t == CfgTypeJson {
		err = jsoniter.Unmarshal(cfg, data)
		if err != nil {
			plog.Error("json Unmarshal file error", zap.String("path", path), zap.Error(err))
		}
	}

	if err != nil {
		return err
	}

	handle.Lock()
	defer handle.Unlock()

	handle.cache[key] = data

	return nil
}

func Get(key string) interface{} {
	handle.RLock()
	defer handle.RUnlock()
	if _, ok := handle.cache[key]; ok {
		return handle.cache[key]
	}

	return nil
}
