package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string    `gorm:"type:text"`
	Weight     float32   `gorm:"type:float"`
	Price      float32   `gorm:"type:float"`
	CreateTime time.Time `gorm:"column:createTime;default:null;type:timestamp" json:"createTime"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("ibothub.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Name:       "香蕉",
		Weight:     1.0,
		Price:      2.8,
		CreateTime: time.Now(),
	})

	var productList []Product

	db.Find(&productList)

	for _, product := range productList {
		fmt.Printf("%v\n", product)
	}
}
