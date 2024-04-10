package pcache

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"github.com/coocood/freecache"
	"sync"
)

type CacheUnit int

const (
	B CacheUnit = iota
	K
	M
	G
)

type Service struct {
}

type CacheItem struct {
	name   string
	handle *freecache.Cache
}

var pool sync.Map

func Alloc(cacheName string, size int, unit CacheUnit) *CacheItem {
	if value, ok := pool.Load(cacheName); ok {
		return value.(*CacheItem)
	}
	switch unit {
	case B:
	case K:
		size *= 1024
	case M:
		size *= 1048576 //1024 * 1024
	case G:
		size *= 1073741824 //1024 * 1024 * 1024
	}
	alloc := freecache.NewCache(size)
	item := &CacheItem{
		name:   cacheName,
		handle: alloc,
	}
	pool.Store(cacheName, item)

	return item
}

func (c *CacheItem) Set(key string, value string, expire int) error {
	return c.handle.Set([]byte(key), []byte(value), expire)
}

func (c *CacheItem) Get(key string) ([]byte, error) {
	return c.handle.Get([]byte(key))
}

func (c *CacheItem) SetString(key string, value string, expire int) error {
	return c.handle.Set([]byte(key), []byte(value), expire)
}

func (c *CacheItem) GetString(key string) (string, error) {
	var value string = ""
	bv, err := c.handle.Get([]byte(key))
	if err != nil {
		return value, err
	}

	return string(bv), nil
}

func (c *CacheItem) SetInt(key string, value int64, expire int) error {
	byteBuf := bytes.NewBuffer([]byte{})
	err := binary.Write(byteBuf, binary.BigEndian, value)
	if err != nil {
		return err
	}

	return c.handle.Set([]byte(key), byteBuf.Bytes(), expire)
}

func (c *CacheItem) GetInt(key string) (int64, error) {
	value, err := c.handle.Get([]byte(key))
	if err != nil {
		return 0, err
	}

	var v int64 = 0
	bufByte := bytes.NewBuffer(value)
	err = binary.Read(bufByte, binary.BigEndian, &v)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (c *CacheItem) SetAny(key string, value interface{}, expire int) error {
	byteBuf := bytes.NewBuffer([]byte{})
	enc := gob.NewEncoder(byteBuf)
	err := enc.Encode(value)
	if err != nil {
		return err
	}

	return c.handle.Set([]byte(key), byteBuf.Bytes(), expire)
}

func (c *CacheItem) GetAny(key string, out interface{}) error {
	value, err := c.handle.Get([]byte(key))
	if err != nil {
		return err
	}

	bufByte := bytes.NewBuffer(value)
	dec := gob.NewDecoder(bufByte)
	err = dec.Decode(out)
	if err != nil {
		return err
	}

	return nil
}
