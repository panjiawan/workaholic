package putils

import (
	"fmt"
	"strconv"
	"strings"
)

func SerializeIntToStr(data []int, sep string) string {
	s := make([]string, 0, len(data))
	for _, v := range data {
		s = append(s, fmt.Sprintf("%d", v))
	}

	return strings.Join(s, sep)
}

func UnSerializeStrToInt(data string, sep string) []int {
	strs := strings.Split(data, sep)
	s := make([]int, 0, len(strs))
	for _, v := range strs {
		if d, err := strconv.Atoi(v); err == nil {
			s = append(s, d)
		}
	}

	return s
}

func SerializeInt32ToStr(data []int32, sep string) string {
	s := make([]string, 0, len(data))
	for _, v := range data {
		s = append(s, fmt.Sprintf("%d", v))
	}

	return strings.Join(s, sep)
}

func UnSerializeStrToInt32(data string, sep string) []int32 {
	strs := strings.Split(data, sep)
	s := make([]int32, 0, len(strs))
	for _, v := range strs {
		if d, err := strconv.Atoi(v); err == nil {
			s = append(s, int32(d))
		}
	}

	return s
}

func SerializeInt64ToStr(data []int64, sep string) string {
	s := make([]string, 0, len(data))
	for _, v := range data {
		s = append(s, fmt.Sprintf("%d", v))
	}

	return strings.Join(s, sep)
}

func UnSerializeStrToInt64(data string, sep string) []int64 {
	strs := strings.Split(data, sep)
	s := make([]int64, 0, len(strs))
	for _, v := range strs {
		if d, err := strconv.Atoi(v); err == nil {
			s = append(s, int64(d))
		}
	}

	return s
}
