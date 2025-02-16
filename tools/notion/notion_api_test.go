package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jomei/notionapi"
	"testing"
)

func TestNotionApi(t *testing.T) {
	client := notionapi.NewClient("secret_70IetZCR48v5IxPglNa9bSySl1ESvLJc6QbNUJqB3Bs")
	me, _ := client.User.Me(context.Background())
	fmt.Println(me.Name)
	page, err := client.Page.Get(context.Background(), "your_page_id")
	if err != nil {
		// 处理错误
	}

	bytes, _ := json.Marshal(page)
	fmt.Println(string(bytes))

}
