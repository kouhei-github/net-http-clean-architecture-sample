package route

import (
	"net-http/myapp/controller"
	"net-http/myapp/controller/auth_handle"
	"net-http/myapp/repository"
	"net-http/myapp/usecase/auth_usecase"
	"net/http"
)

func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// 練習
	mux.HandleFunc("/", controller.Handler)
	mux.HandleFunc("/two", controller.HandlerTwo)

	// ユーザー登録
	signup := auth_handle.NewSignupHandler(&auth_usecase.Signup{AdminUserRepo: &repository.Administer{}})
	mux.HandleFunc("/api/v1/signup", signup.SignupHandler)

	return mux
}
