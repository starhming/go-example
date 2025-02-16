package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestStringBase(t *testing.T) {
	user := User{
		UserId:   "123",
		UserName: "test",
	}
	err := redisCli.Set(context.Background(), "user_key_test", &user, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	res := User{}
	err = redisCli.Get(context.Background(), "user_key_test").Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
