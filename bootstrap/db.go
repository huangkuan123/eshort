package bootstrap

import (
	"eshort/pkg/config"
	"eshort/pkg/model"
	"fmt"
	"time"
)

func SetUpDB() {
	db := model.ConnectDB()
	//database/sql 包里的 *sql.DB 对象
	sqlDB, _ := db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	fmt.Println("数据库初始化成功")
}
