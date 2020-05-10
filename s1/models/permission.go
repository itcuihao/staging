package models

import "time"

type Permission struct {
	Id          int       `gorm:"type:int unsigned auto_increment;primary_key"`
	Title       string    `gorm:"type:varchar(50);default:'';not null;comment:'权限'"`
	Description string    `gorm:"type:varchar(50);default:'';not null;comment:'描述'"`
	Seq         int       `gorm:"type:int;default:0;not null;comment:'排序'"`
	Status      int       `gorm:"type:tinyint;default:0;not null;comment:'状态'"`
	CreateTime  time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime  time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}

func (Permission) TableName() string {
	return "permissions"
}
