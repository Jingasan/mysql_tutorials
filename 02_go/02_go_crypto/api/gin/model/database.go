package model

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	database_name := os.Getenv("MYSQL_DATABASE")        // データベース名
	root_password := os.Getenv("MYSQL_ROOT_PASSWORD")   // ROOTパスワード
	db_container_ipv4 := os.Getenv("DB_CONTAINER_IPV4") // DBコンテナIPv4
	db_container_port := os.Getenv("DB_CONTAINER_PORT") // DBコンテナポート番号
	var path string = "root:" + root_password + "@tcp(" + db_container_ipv4 + ":" + db_container_port + ")/" + database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dialector := mysql.Open(path)
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}
