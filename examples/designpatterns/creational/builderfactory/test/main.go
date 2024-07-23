package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateMACAddress() string {
	// 创建一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 本地管理的MAC地址第二个字节的第二位应该为1
	// 这里我们使用微软的MAC地址前缀"00-15-5D"作为示例
	mac := []byte{0x00, 0x15, 0x5D, 0x00, 0x00, 0x00}

	// 生成后三个字节的随机值
	for i := 3; i < 6; i++ {
		mac[i] = byte(r.Intn(256))
	}

	// 格式化MAC地址字符串
	return fmt.Sprintf("%02X-%02X-%02X-%02X-%02X-%02X",
		mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

func main() {
	macAddress := generateMACAddress()
	fmt.Println("Generated MAC Address:", macAddress)
}
