package models

import "time"

type Role struct {
	Id          int       `gorm:"type:int unsigned auto_increment;primary_key"`
	Title       string    `gorm:"type:varchar(50);default:'';not null;comment:'角色'"`
	Description string    `gorm:"type:varchar(50);default:'';not null;comment:'描述'"`
	Seq         int       `gorm:"type:int;default:0;not null;comment:'排序'"`
	Status      int       `gorm:"type:tinyint;default:0;not null;comment:'状态'"`
	CreateTime  time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime  time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}

func (Role) TableName() string {
	return "roles"
}

type RolePermission struct {
	RoleId       int       `gorm:"type:int unsigned;primary_key;auto_increment:false"`
	PermissionId int       `gorm:"type:int unsigned;primary_key;auto_increment:false"`
	Status       int       `gorm:"type:tinyint;default:0;not null;comment:'状态'"`
	CreateTime   time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime   time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
