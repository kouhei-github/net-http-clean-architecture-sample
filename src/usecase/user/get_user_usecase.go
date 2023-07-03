package user

import (
	"fmt"
	"net-http/myapp/domain"
)

type GetUser struct {
	AdminUserRepo domain.AdminUserRepository
	JwtRepo       domain.AuthJwtToken
}

func (usr *GetUser) GetUserData(jwtToken string) (float64, error) {

	userId, err := usr.JwtRepo.AuthorizationProcess(jwtToken)
	fmt.Println(userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
