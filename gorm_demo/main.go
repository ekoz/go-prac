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

	// if len(name) > 0 {
	// 	db.Find(&userList, "name like ?", name)
	// } else {
	// 	db.Find(&userList)
	// }

	db.Where("name like ?", name).Find(&userList)

	for _, user := range userList {
		fmt.Printf("%v\n", user)
	}
}

// 更新
func update(db *gorm.DB) {
	var u1, u2 User

	db.First(&u1, 1)
	fmt.Printf("%v\n", u1)
	u1.Age = u1.Age + 1
	// save 是更新所有字段
	db.Debug().Save(u1)

	db.First(&u2, 2)
	fmt.Printf("%v\n", u2)
	// updates 或 update 是更新指定字段，而且是该字段值存在变化时才更新
	// update 只更新一个字段
	db.Debug().Model(&u2).Updates(User{Age: 20})

	m1 := map[string]any{
		"name": "zhangfei",
		"age":  32,
		"sex":  1,
	}
	db.Debug().Model(&u2).Updates(m1)

	// 忽略某些字段，更新其他字段
	// db.Model(&u2).Omit("name", "age").Updates(m1)

	// 选定更新字段，只更新 age 字段
	// db.Model(&u2).Select("age").Updates(m1)

	// 让所有数据年龄都加上1
	db.Debug().Model(&User{}).Where("age>1").Updates(map[string]interface{}{"age": gorm.Expr("age*?+?", 1, 1)})

}

// 删除
func delete(db *gorm.DB) {
	// 必须保证主键有值，如果主键没有值，相当于是删除所有数据
	// 逻辑删除
	var u1 = User{}
	u1.ID = 1
	db.Delete(&u1)

	db.Where("name like ?", "%司马%").Delete(&User{})

	// 物理删除
	db.Unscoped().Delete(&u1)

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
	// findOne(db, 4)

	//
	// findByName(db, "司马%")

	// update(db)

	delete(db)

}
