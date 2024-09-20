package domain

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        string         // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     time.Time      // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Name == "Jinzhu" {
		return errors.New("invalid user")
	}
	return
}
