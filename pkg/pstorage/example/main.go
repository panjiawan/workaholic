package main

import (
	"bytes"
	"fmt"

	"github.com/panjiawan/workaholic/pkg/pstorage"
)

func main() {
	s := pstorage.New(pstorage.Qiniu, &pstorage.Config{
		SecretId:      "",
		SecretKey:     "",
		UseHTTPS:      false,
		UseCdnDomains: false,
	})
	b := []byte("hello, this is qiniu cloud")
	bio := bytes.NewReader(b)
	res, err := s.PutFromStream("test_bucket", "avatar/test/test.log", bio, int64(len(b)))
	fmt.Println(res, err)
}
