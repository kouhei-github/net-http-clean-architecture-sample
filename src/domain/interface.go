package domain

import (
	"net-http/myapp/domain/model/user"
)

type AdminUserRepository interface {
	SaveAdminUser(user *user.AdminUser) error
	FindAdminUserByEmail(email string) (*user.AdminUser, error)
}
