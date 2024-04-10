package pcache

import "testing"

type Item struct {
	G int
}

type TestStruct struct {
	A int
	B string
	C *Item
}

func TestCache(t *testing.T) {
	item := &TestStruct{
		A: 1,
		B: "hello world",
		C: &Item{
			G: 3,
		},
	}
	itemCache := Alloc("test", 1, B)
	err := itemCache.SetAny("tk", item, 0)
	t.Logf("%v,%v", item.A, err)

	getItem := &TestStruct{}
	err = itemCache.GetAny("tk", getItem)
	t.Logf("%v,%v", getItem.C.G, err)
	itemCache.SetInt("ti", 1, 0)
	i, err := itemCache.GetInt("t2i")
	t.Logf("%v,%v", i, err)

}
