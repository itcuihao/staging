package dao

import "github.com/itcuihao/staging/s1/models"

func (dao *Dao) GetUserById(id string) (*models.User, error) {
	user := &models.User{}
	if err := dao.db.Where("id=?", id).FirstOrInit(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *Dao) CreateOrFirstUserByAccount(account string) (*models.User, error) {
	var user = &models.User{
		Account: account,
	}

	if err := dao.db.Where("account = ? AND status >= 0", account).FirstOrCreate(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (dao *Dao) UpdateUserToken(id int, token string, expireAt int64) error {
	return dao.db.Model(models.User{}).Where("id=?", id).Updates(models.User{
		AccessToken:   token,
		TokenExpireAt: expireAt,
	}).Error
}

func (dao *Dao) GetUserRoleIds(id int) ([]int, error) {
	roleIds := make([]int, 0)
	if err := dao.db.Model(models.UserRole{}).Where("user_id=?", id).Pluck("role_id", &roleIds).Error; err != nil {
		return nil, err
	}
	return roleIds, nil
}

func (dao *Dao) ExistUserRole(userId, roleId int) (bool, error) {
	var count int
	if err := dao.db.Model(models.UserRole{}).Where("user_id=? AND role_id=?", userId, roleId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
