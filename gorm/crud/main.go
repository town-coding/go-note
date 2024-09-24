package main

import (
	"fmt"
	"go-note/gorm/domain"
	"go-note/gorm/initialize"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func main() {
	db = initialize.InitMysql()
	//Create()

	//Select()

	Update()

}

func Update() {
	var user domain.User
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	// 会保存所有的字段，即使字段是零值
	db.Save(&user)

	// Save() 保存时不包含主键，他将执行 Create, 否者执行 Update
	db.Save(&domain.User{Name: "qwe", Age: 124})

	// 更新单个列，根据条件更新
	db.Model(&domain.User{}).Where("name = ?", "qwe").Update("age", gorm.Expr("age - 10"))
}

func Select() {
	var user domain.User
	// 获取第一条记录（主键升序）
	//result := db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	//result := db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	//result := db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	//fmt.Printf("result: %v", result)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	//fmt.Printf("err record not found %v", errors.Is(result.Error, gorm.ErrRecordNotFound))

	//var users []domain.User

	// works because destination struct is passed in
	//db.Session(&gorm.Session{QueryFields: true}).First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// works because model is specified using `db.Model()`
	result := map[string]interface{}{}
	//db.Model(&domain.User{}).First(&result)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	//result := map[string]interface{}{}
	//tx := db.Table("users").First(&result)

	// works with Take
	//result := map[string]interface{}{}
	//tx := db.Session(&gorm.Session{QueryFields: true}).Table("users").Take(&result)
	tx := db.Model(&domain.User{}).Where("id > (?)", db.Model(&domain.User{}).Select("AVG(id)")).Find(&result)

	fmt.Printf("err: %v", tx.Error)

	db.Where(domain.User{Name: "Jinzhu"}).Assign(domain.User{Age: 10}).FirstOrInit(&user)
	fmt.Printf("user: %v", user)

}

func Create() {
	//err := db.AutoMigrate(&domain.User{})
	//if err != nil {
	//	panic(err)
	//}
	//user := domain.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Create(&user)

	users := []*domain.User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	//err := db.CreateInBatches(&users, len(users)).Error
	err := db.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(&users, len(users)).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
