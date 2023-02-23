package mysql

import (
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Id       int
	Username string
	Age      int
}

func TestInitMysql(t *testing.T) {
	InitMysql()
	user := User{
		Username: "Lee",
		Age:      18,
	}
	Engine.Create(&user)
	user2 := User{
		Id:       3,
		Username: "Lee",
		Age:      18,
	}
	Engine.Delete(&user2)
}
