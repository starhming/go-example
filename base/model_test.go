package base

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/starshm/go-example/model/visibility"
)

func TestStructVisibility(t *testing.T) {
	data := `{"window_size": 60, "count": 5, "continuous": true}`
	orc := visibility.OperatorRuleConfig{}
	err := json.Unmarshal([]byte(data), &orc)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(orc.GetCount())
	fmt.Println(orc.GetCount())
	fmt.Println(orc.GetContinuous())

}
