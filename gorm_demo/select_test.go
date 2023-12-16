package gorm_demo

import (
	"fmt"
	"testing"

	"github.com/starshm/go-example/model"
)

func TestSelectGroup(t *testing.T) {
	u := model.User{}
	result := db.Where(db.Where("name = ?", "hm")).
		Where(db.Where("age = ?", 18).Or("age = ?", 20)).
		First(&u)
	if result != nil {

	}
	fmt.Println(u)
}
