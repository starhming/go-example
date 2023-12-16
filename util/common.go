package util

import (
	"encoding/json"
	"fmt"
)

func PrintJsonStruct(t any) {
	marshal, _ := json.Marshal(t)
	fmt.Println(string(marshal))
}
