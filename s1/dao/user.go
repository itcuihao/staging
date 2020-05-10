package dao

import "github.com/itcuihao/staging/s1/models"

func (db *Dao) GetUserById(id string) (*models.User, error) {
	user := &models.User{}
	if err := db.db.Where("id=?", id).FirstOrInit(user).Error; err != nil {
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

func (db *Dao) GetUserRoleIds(id int) ([]int, error) {
	roleIds := make([]int, 0)
	if err := db.db.Where("user_id=?", id).Pluck("role_id", &roleIds).Error; err != nil {
		return nil, err
	}
	return roleIds, nil
}
