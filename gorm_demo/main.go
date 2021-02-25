package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*var mysqlConfig = config.Config.Mysql
	var ptr = &mysqlConfig.Username
	*ptr = "hahaha"
	fmt.Printf("Mysql addr --> %p\n", &(config.Config.Mysql.Username))
	fmt.Printf("mysqlConfig addr --> %p\n", &(mysqlConfig.Username))
	fmt.Println(config.Config.Mysql)
	dsn := "root:abc123@tcp(192.168.131.125:3306)/iplaymtg?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Create()*/

}

type user struct {

}
