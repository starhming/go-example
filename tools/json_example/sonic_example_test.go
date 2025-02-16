package json_example

import (
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

func TestSonicMashalUnMashl(t *testing.T) {
	type Person struct {
		Age  int64
		Name string
	}
	p := Person{Age: 60, Name: "tom"}

	bytes, err := sonic.Marshal(p)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(bytes))

	a := getBytesArr()
	fmt.Println(string(a))
}

func getBytesArr() []byte {
	return nil
}

func TestSonicGet(t *testing.T) {
	// 示例 JSON 字符串
	jsonStr := `{
		"person": {
			"name": "Alice",
			"age": 30,
			"address": {
				"city": "New York",
				"postalCode": "10001"
			}
		}
	}`
	fromString, err2 := sonic.GetFromString(jsonStr, "person", "name")
	if err2 != nil {
		fmt.Println(fromString.Exists())
	}

	path, err := keyExistsAtPath(jsonStr, "person.name")
	fmt.Println(path, "  ", err)

	path, err = keyExistsAtPath(jsonStr, "person.name.age")
	fmt.Println(path, "  ", err)

}

func keyExistsAtPath(jsonStr string, path string) (bool, error) {
	// 解析 JSON
	node, err := sonic.Get([]byte(jsonStr), path)
	if err != nil {
		return false, err
	}

	// 检查是否是空值，或者该路径不存在
	if node.Exists() {
		return true, nil
	}

	return true, nil
}

func TestJsonDeleteKey(t *testing.T) {
	// 示例 JSON 字符串
	jsonStr := `{
		"person": {
			"name": "Alice",
			"age": 30,
			"address": {
				"city": "New York",
				"postalCode": "10001"
			}
		}
	}`

	// 要删除的路径，例如删除 "person.address.city"
	path := "person.address.city"

	// 调用函数删除指定路径下的 key
	modifiedJSON, err := removeKeyAtPath(jsonStr, path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印修改后的 JSON 字符串
	fmt.Println("Modified JSON:", modifiedJSON)
}

func TestExistsKeyInJSON(t *testing.T) {
	// 示例 JSON 字符串
	jsonStr := `{
		"person": {
			"name": "Alice",
			"age": 30,
			"address": {
				"city": "New York",
				"postalCode": "10001"
			}
		}
	}`

	// 要删除的路径，例如删除 "person.address.city"
	path := "person.address.city"

	fmt.Println(ExistsKeyInJSON(jsonStr, path))
}

func TestExistsKeyInJSON1(t *testing.T) {
	// 示例 JSON 字符串
	jsonStr := `{"number":123,"object":{"a":"b","c":"d","e":{"f":"g","h":"i"}}}`

	// 要删除的路径，例如删除 "person.address.city"
	path := "abject.e.f"

	//inputJSON: `{"number":123,"object":{"a":"b","c":"d","e":{"f":"g","h":"i"}}}`, keyPath: "abject.e.f"

	fmt.Println(ExistsKeyInJSON(jsonStr, path))
}

func TestUmarshalArr(t *testing.T) {
	data := `[1,2,3,4,5,6]`
	arr := make([]int, 10, 10)
	err := sonic.UnmarshalString(data, &arr)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(arr)
}
