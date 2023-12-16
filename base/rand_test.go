package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandSeed(t *testing.T) {
	// 使用当前时间的纳秒数来初始化随机数生成器的种子
	rand.Seed(time.Now().UnixNano())

	// 生成并打印一个随机数
	随机数 := rand.Intn(100)
	fmt.Println("随机数:", 随机数)
}
