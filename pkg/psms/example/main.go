package main

import (
	"fmt"
	"github.com/panjiawan/workaholic/pkg/psms"
)

func main() {
	s := psms.New(psms.QCloud, &psms.Config{
		SecretId:   "",
		SecretKey:  "",
		SdkAppId:   "",
		Sign:       "",
		TemplateId: "",
	})
	res, err := s.Send("+8613500000000", []string{"6666"})
	fmt.Println(res, err)
}
