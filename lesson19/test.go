package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func test() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	// 迁移 schema
	db.AutoMigrate(&Product{})
	// create
	db.Create(&Product{Code: "D402", Price: 100})
	// Read
	var product Product
	db.First(&product, 1) // 根据整型主见查找
	db.First(&product, "code = ?", "D42") // 查找code字段值为D42的记录

	// Update   - 将product的price更新为200
	db.Model(&product).Update("Price", 200)
	// Update -- 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	db.Delete(&product, 1)
}


type Product struct {
	gorm.Model
	Code string
	Price uint
}