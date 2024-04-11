package putils

import "encoding/base64"

// BASE64
func Base64Encode(bytes []byte) string {
	str := base64.StdEncoding.EncodeToString(bytes)
	return str
}

func Base64Decode(data string) []byte {
	decodeData, _ := base64.StdEncoding.DecodeString(data)
	return decodeData
}
