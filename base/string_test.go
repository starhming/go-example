package base

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestStringBaseTest(t *testing.T) {
	streamName := " stream-691400032724189962  "
	streamName = strings.TrimSpace(streamName)
	fmt.Println(streamName)
}

func TestJsonMarshal(t *testing.T) {
	a := 1
	fmt.Println(MarshalStruct(a))
}

func MarshalStruct(a any) string {
	marshal, err := json.Marshal(a)
	if err != nil {
		fmt.Println("marshal failed: ", err.Error())
		return ""
	}

	return string(marshal)
}
