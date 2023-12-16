package gorm_demo

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/starshm/go-example/model"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	now := time.Now()
	user := model.User{Name: "hm", Birthday: &now}

	result := db.Create(&user) // 通过数据的指针来创建

	//user.ID             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数

	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func TestAutoMigrate(t *testing.T) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	user := &model.User{}
	db.First(user)

	db.Model(&model.User{}).Where("age = ?", 0).Update("name", "hello")
}

func TestFirstOrCreate(t *testing.T) {
	now := time.Now()
	user := model.User{Name: "minghai", Birthday: &now}

	result := db.Where("name = ?", "minghai").FirstOrCreate(&user)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	fmt.Println(user)
}
