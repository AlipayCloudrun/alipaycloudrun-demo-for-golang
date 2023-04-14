package app

import (
	"alipaycloudrun-demo-for-go/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() *gorm.DB {
	dsn := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// 从环境变量中获取数据库连接参数，如果获取不到则会使用默认值
	// 用户名
	user := util.GetEnvDefault("DATABASE_USERNAME", "root")
	// 密码
	pwd := util.GetEnvDefault("DATABASE_PASSWORD", "123456")
	// 数据库连接地址及端口
	addr := util.GetEnvDefault("DATABASE_HOST", "127.0.0.1:3306")
	// 数据库名
	dataBase := util.GetEnvDefault("DATABASE_NAME", "test")

	dsn = fmt.Sprintf(dsn, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		fmt.Printf("Db open error, err:%+v", err.Error())
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("DB init error, err:%+v", err.Error())
		return nil
	}

	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("finish init mysql with ", dsn)
	return db
}
