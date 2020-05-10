package tool

import (
	"fmt"
	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"
	"github.com/itcuihao/staging/s1/models"
)

func AutoMigrate(dao *dao.Dao) {
	db := dao.NewDB()
	tableEngine := `ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '%s'`
	exist := db.HasTable(&models.User{})
	common.Log.Info(exist)
	if !exist {
		db.Set("gorm:table_options", fmt.Sprintf(tableEngine, "用户表")).CreateTable(&models.User{})
	}
	db.Set("gorm:table_options", fmt.Sprintf(tableEngine, "角色表")).CreateTable(&models.Role{})
	db.Set("gorm:table_options", fmt.Sprintf(tableEngine, "权限表")).CreateTable(&models.Permission{})
	db.Set("gorm:table_options", fmt.Sprintf(tableEngine, "用户角色表")).CreateTable(&models.UserRole{})
	db.Set("gorm:table_options", fmt.Sprintf(tableEngine, "角色权限表")).CreateTable(&models.RolePermission{})
}
