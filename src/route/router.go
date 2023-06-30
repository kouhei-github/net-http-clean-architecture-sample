package route

import (
	"net-http/myapp/controller"
	"net-http/myapp/controller/auth_handle"
	"net-http/myapp/repository"
	"net-http/myapp/usecase/auth_usecase"
	"net/http"
)

type Router struct {
	Mutex *http.ServeMux
}

func (router *Router) GetRouter() {

	// 練習
	router.Mutex.HandleFunc("/", controller.Handler)
	router.Mutex.HandleFunc("/two", controller.HandlerTwo)

	// ユーザー登録
	signup := auth_handle.NewSignupHandler(&auth_usecase.Signup{AdminUserRepo: &repository.Administer{}})
	router.Mutex.HandleFunc("/api/v1/signup", signup.SignupHandler)
}
