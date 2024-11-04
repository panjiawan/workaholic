package putils

import "testing"

func TestSerialize(t *testing.T) {
	a := SerializeInt64ToStr([]int64{1, 2, 3}, ",")
	t.Log(a)

	b := UnSerializeStrToInt64(a, ",")
	t.Log(b)
}
