package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study.go/gorm_demo/config"
)

func main() {
	var mysqlConfig = config.Config.Mysql
	var ptr = &mysqlConfig.Username
	*ptr = "hahaha"
	fmt.Printf("Mysql addr --> %p\n", &(config.Config.Mysql.Username))
	fmt.Printf("mysqlConfig addr --> %p\n", &(mysqlConfig.Username))
	fmt.Println(config.Config.Mysql)
	dsn := "root:abc123@tcp(192.168.131.125:3306)/iplaymtg?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	/*var u user
	db.Raw("select id, username, head, fire from user where id = ?", 1453295).Scan(&u)
	fmt.Println(u)
	var demo = Demo{Name: "falkj"}
	db.Create(&demo)*/
	/*total := 0
	res := db.Raw("select count(1) from user").Scan(&total)
	fmt.Println("total -->", total, ",err --> ", res.Error)*/
	var u user
	var u1 user
	db.Find(&u1)
	//db.Exec("select * from user where username = ?", "唐青枫本人").Find(&u)
	fmt.Println(u)
	fmt.Println(u1)

}

type user struct {
	Id       int
	Username string
	Head     string
	Fire     uint
}
func (u *user) TableName () string {
	return "user"
}

type Demo struct {
	gorm.Model
	Name string `gorm:"comment:'名字'"`
	Age  uint   `gorm:"comment:'年龄' default:'20'"`
}
