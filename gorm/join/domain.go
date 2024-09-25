package main

import "gorm.io/gorm"

//type User struct {
//	gorm.Model
//	Name      string
//	CompanyID int
//	Company   Company // 使用 Company.CompanyID 作为引用
//}

type Company struct {
	CompanyID int `gorm:"primaryKey;index"` // 定义为主键
	Code      string
	Name      string
}

func (User) TableName() string {
	return "sys_user"
}

func (Company) TableName() string {
	return "sys_company"
}

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserID"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

// User 拥有并属于多种 language，`user_languages` 是连接表
//type User struct {
//	gorm.Model
//	Languages []*Language `gorm:"many2many:user_languages;"`
//}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}
