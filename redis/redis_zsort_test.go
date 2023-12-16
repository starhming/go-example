package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestZsortBase(t *testing.T) {
	zRange, err := RedisZRange[string](context.Background(), "student", 0, 0)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, a := range zRange {
		fmt.Println(a)
	}
}

func RedisZRange[T any](ctx context.Context, key string, start, end int64) ([]T, error) {
	arr, err := rdb.ZRangeWithScores(ctx, key, start, end).Result()
	if err != nil {
		return nil, err
	}
	var res []T
	for _, a := range arr {
		s := a.Member.(string)
		var t T
		if isString(t) {
			t = any(s).(T)
		} else {
			if err = json.Unmarshal([]byte(s), &t); err != nil {
				fmt.Printf("RedisZPopMin Unmarshal member failed, err: %s", err.Error())
				continue
			}
		}
		if v, ok := any(t).(SetScorer); ok {
			v.SetScore(a.Score)
		}
		res = append(res, t)
	}
	return res, err

}
func isInt[T any](x T) (ok bool) {

	_, ok = any(x).(int) // convert, then assert
	return
}

func isString[T any](x T) (ok bool) {

	_, ok = any(x).(string) // convert, then assert
	return
}
