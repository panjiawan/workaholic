package putils

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net"
	"time"
)

func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// Rand [min, max]
func Rand(min int32, max int32) int32 {
	n := max - min + 1
	randNum := rand.Int31n(n) + min

	return randNum
}

// [min, max)
func RealRand(min, max int64) int64 {
	num := max - min
	randNum, _ := crand.Int(crand.Reader, big.NewInt(num))
	return randNum.Int64()
}

// BytesToInt32Array bytes转为int32数组
func BytesToInt32Array(byteData []byte) []uint32 {
	byteLen := len(byteData)
	int32Array := make([]uint32, 0, 1024)
	max := int(math.Ceil(float64(byteLen) / 4))
	for i := 0; i < max; i++ {
		begin := i * 4
		end := begin + 4
		if end > byteLen {
			end = byteLen
		}
		soleNum := binary.LittleEndian.Uint32(byteData[begin:end])
		int32Array = append(int32Array, soleNum)
	}

	return int32Array
}

// GenRandomNumber 生成count个start到end之间的随机数[start,end)
func GenRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

// 获取IP地址
func LocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ip string = "localhost"
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}

// 乱序一个数组的key
func ConfusionArrayKey(arrayLen int) []int {
	array := make([]int, 0)
	for i := 0; i < arrayLen; i++ {
		array = append(array, 0)
	}

	rand.Seed(time.Now().UnixNano())
	keys := make([]int, 0)
	for i := 0; i < arrayLen; i++ {
		tLen := len(array)
		index := rand.Intn(tLen)
		array = append(array[0:index], array[index+1:]...)
		keys = append(keys, index)
	}

	return keys
}
