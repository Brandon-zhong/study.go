package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study.go/gorm_demo/config"
)

func main() {
	var admin = config.Config.Mysql
	var dsn = admin.Username + ":" + admin.Password + "@(" + admin.Url + ")/" + admin.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var u user
	db.Raw("select id, username, head, fire from user where id = ?", 1453295).Scan(&u)
	fmt.Println(u)
	var demo = Demo{Name: "falkj"}
	db.Create(&demo)

}

type user struct {
	Id       int
	Username string
	Head     string
	Fire     uint
}

type Demo struct {
	gorm.Model
	Name string `gorm:"comment:'名字'"`
	Age  uint   `gorm:"comment:'年龄' default:'20'"`
}
