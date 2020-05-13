package dao

import "github.com/itcuihao/staging/s1/models"

func (dao *Dao) GetRoleByIds(ids ...int) ([]*models.Role, error) {
	roles := make([]*models.Role, 0, len(ids))
	if err := dao.db.Where("id in (?)", ids).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (dao *Dao) GetRoleByTitle(title string) (*models.Role, error) {
	role := &models.Role{}
	if err := dao.db.Where("title=?", title).FirstOrInit(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
