package putils

import (
	"bytes"
	"encoding/binary"
)

// Int64ToBytes 字节与int互换--大端法
func Int64ToBytes(n int64) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// BytesToInt64 字节与int互换--大端法
func BytesToInt64(bys []byte) int64 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}

// Int32ToBytes 字节与int互换--大端法
func Int32ToBytes(n int32) []byte {
	data := n
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// BytesToInt32 字节与int互换--大端法
func BytesToInt32(bys []byte) int32 {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}

// 字节与int互换--大端法
func UInt16ToBytes(n int) []byte {
	data := uint16(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// 字节与int互换--大端法
func BytesToUInt16(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data uint16
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

//func ByteSlice2string(b []byte) string {
//	return unsafe.String(&b[0], len(b))
//}
//
//func String2ByteSlice(s string) []byte {
//	return unsafe.Slice(unsafe.StringData(s), len(s))
//}
