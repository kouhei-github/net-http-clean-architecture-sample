package repository

import (
	"net-http/myapp/domain/model/user"
)

type Administer struct {
}

func (r *Administer) SaveAdminUser(user *user.AdminUser) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Administer) FindAdminUserByEmail(email string) (*user.AdminUser, error) {
	var adminUser user.AdminUser
	result := db.Where("email = ?", email).First(&adminUser)
	if result.Error != nil {
		return &adminUser, result.Error
	}
	return &adminUser, nil
}
