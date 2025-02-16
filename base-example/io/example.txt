package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadFileByLine(t *testing.T) {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建一个Scanner来按行读取文件
	scanner := bufio.NewScanner(file)

	// 按行扫描文件
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行内容
		fmt.Println(line)      // 输出当前行
	}

	// 检查是否有扫描错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("扫描文件时出错: %v", err)
	}
}
