package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255)"`
	Age      int       `gorm:"type:int(3)"`
	Sex      int       `gorm:"type:int(2)"`
	Birthday time.Time `gorm:"column:birthday;default:null" json:"birthday"`
}

// 自定义 user 表表名
func (User) TableName() string {
	return "t_user"
}

// 自定义 动态表名
func UserTable(u *User) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("t_user_" + u.Name)
	}
}

// 插入数据
func insert(db *gorm.DB) {
	db.Create(&User{
		Name: "小明",
		Age:  18,
		Sex:  1,
	})
}

// 批量插入数据
func batchInsert(db *gorm.DB) {
	var userList []User

	userList = append(userList, User{
		Name: "小红",
		Age:  16,
		Sex:  0,
	}, User{Name: "司马光", Age: 8, Sex: 1})

	db.Create(&userList)

	for _, user := range userList {
		fmt.Printf("%v\n", user)
	}
}

// 查询数据
func findOne(db *gorm.DB, id int) {
	u := User{}
	db.First(&u, id)
	// db.First(&u, "id = ?", "1")
	fmt.Printf("%v\n", u)
}

// 查询数据
func findByName(db *gorm.DB, name string) {
	var userList []User
	if len(name) > 0 {
		db.Find(&userList, "name like ?", name)
	} else {
		db.Find(&userList)
	}
	for _, user := range userList {
		fmt.Printf("%v\n", user)
	}
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/auth2?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 表名前缀和是否加s规范
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "t_",
	// 		SingularTable: true,
	// 	},
	// 	Logger: logger.Default.LogMode(logger.Silent),
	// })

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// 这一步，会自动创建表，orm框架默认的表名是加 s，比如 User struct 创建的表就是 users
	db.AutoMigrate(&User{})

	// 动态表名
	// db.Scopes(UserTable(&user)).Create(&user)

	// 自定义表名
	// db.Table("user").AutoMigrate(&User{})

	// insert(db)

	// batchInsert(db)

	//
	findOne(db, 4)

	//
	findByName(db, "%司马")

}
