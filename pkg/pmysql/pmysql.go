package pmysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	handler  *gorm.DB
	host     string
	port     int
	username string
	password string
	db       string
	prefix   string
	charset  string
	debug    bool
	maxIdle  int
	maxOpen  int
}

type Option func(*Service)

func New(opts ...Option) *Service {
	s := &Service{}
	for _, f := range opts {
		f(s)
	}

	return s
}

// MysqlInit 初始化
func (s *Service) Run() error {
	if s.charset == "" {
		s.charset = "utf8"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		s.username,
		s.password,
		s.host,
		s.port,
		s.db,
		s.charset)
	handle, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if db, err := handle.DB(); err == nil {
		db.SetMaxIdleConns(s.maxIdle)
		db.SetMaxOpenConns(s.maxOpen)
	} else {
		return err
	}

	if s.debug {
		handle = handle.Debug()
	}

	s.handler = handle

	return nil
}

func (m *Service) Handle() *gorm.DB {
	return m.handler
}

func (m *Service) Close() {
}
