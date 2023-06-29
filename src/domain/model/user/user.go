package user

import "gorm.io/gorm"

type AdminUser struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	CompanyId int    `json:"company_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password"  binding:"required"`
}

func NewAdminUser(name string, companyId int, email string, password string) (*AdminUser, error) {
	return &AdminUser{Email: email, Name: name, CompanyId: companyId, Password: password}, nil
}
