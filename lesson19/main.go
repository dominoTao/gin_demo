package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	//想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数) 要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
	dsn := "root:root@tcp(127.0.0.1:3305)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// 简单连接
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//MySQl 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		//DriverName: "my_mysql_driver", // 自定义驱动名
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})


	// 通过一个现有的数据库连接来初始化 *gorm.DB
	//sqlDB, err := sql.Open("mysql", "mydb_dsn")
	//gorm.Open(mysql.New(mysql.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})


	if err != nil {
		panic(err)
	}
// 创建表 自动迁移 把结构体和表进行对应
	db.AutoMigrate(&UserInfo{})
	// 创建数据对象
	info := UserInfo{1, "小黄", "男", "蛙泳"}
	db.Create(&info)
	//defer
}

// 对应数据表
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

