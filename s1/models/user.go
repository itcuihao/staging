package models

import (
	"time"
)

type User struct {
	Id            int       `gorm:"type:int unsigned auto_increment;primary_key"`
	Account       string    `gorm:"type:varchar(50);default:'';not null;comment:'账号'"`
	Name          string    `gorm:"type:varchar(50);default:'';not null;comment:'姓名'"`
	Status        int       `gorm:"type:tinyint;default:0;not null;comment:'状态'"`
	AccessToken   string    `gorm:"type:varchar(1000);default:'';not null;comment:'token'"`
	TokenExpireAt int64     `gorm:"type:int;default:'';not null;comment:'token过期时间'"`
	CreateTime    time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime    time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}

func (User) TableName() string {
	return "users"
}

type UserRole struct {
	UserId     int       `gorm:"type:int unsigned;primary_key;auto_increment:false"`
	RoleId     int       `gorm:"type:int unsigned;primary_key;auto_increment:false"`
	Status     int       `gorm:"type:tinyint;default:0;not null;comment:'状态'"`
	CreateTime time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}

func (UserRole) TableName() string {
	return "user_role"
}
