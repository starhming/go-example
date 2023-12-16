package pc

import (
	"fmt"
	"testing"
)

/**
在项目根目录下，执行
protoc --proto_path=./ --go_out=./grpc-example/proto-example2 ./grpc-example/proto-example2/*.proto
*/

func TestComputer(t *testing.T) {
	computer := Computer{
		Name: "macbook pro",
		Cpu: &CPU{
			Name:      "Intel Core i9",
			Frequency: "2.4 GHz",
		},
		Memory: &Memory{
			Name: "DDR4",
			Cap:  "32 GB",
		},
	}
	str := computer.String()
	fmt.Println(str)
	// Output: name:"macbook pro" cpu:{Name:"Intel Core i9" Frequency:"2.4 GHz"} memory:{Name:"DDR4" Cap:"32 GB"}
}
