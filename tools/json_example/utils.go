package json_example

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/duke-git/lancet/v2/slice"
	"strings"
)

// 删除指定路径下的 key
func removeKeyAtPath(inputJSON string, path string) (string, error) {
	// 将 JSON 字符串解析为 map
	var data map[string]any
	if err := json.Unmarshal([]byte(inputJSON), &data); err != nil {
		return "", err
	}

	// 将路径按 "." 分割
	keys := strings.Split(path, ".")

	// 递归删除指定路径下的 key
	if err := deleteKey(data, keys); err != nil {
		return "", err
	}

	// 将修改后的 map 重新编码为 JSON 字符串
	modifiedJSON, err := sonic.MarshalString(data)
	if err != nil {
		return "", err
	}

	return modifiedJSON, nil
}

// 递归遍历 JSON 数据结构并删除指定路径的 key
func deleteKey(data map[string]any, keys []string) error {
	// 如果只剩下最后一个 key，直接删除
	if len(keys) == 1 {
		delete(data, keys[0])
		return nil
	}

	// 检查当前 key 对应的值是否是一个嵌套的 map
	if nestedMap, ok := data[keys[0]].(map[string]any); ok {
		// 递归处理嵌套的 map
		return deleteKey(nestedMap, keys[1:])
	}

	return fmt.Errorf("key path not found or not a valid structure")
}

// ExistsKeyInJSON Whether the JSON string contains the specified key path
func ExistsKeyInJSON(inputJSON string, keyPath string) bool {
	if keyPath == "" {
		return false
	}
	splitArr := strings.Split(keyPath, ".")

	keyPathArr := slice.Map(splitArr, func(index int, s string) interface{} {
		return s
	})
	node, err := sonic.GetFromString(inputJSON, keyPathArr...)
	if err != nil {
		return false
	}

	return node.Exists()
}
