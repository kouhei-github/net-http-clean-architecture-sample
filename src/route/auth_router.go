package route

import (
	"net-http/myapp/controller"
)

func (router *Router) GetAuthRouter() {

	// 練習
	router.Mutex.HandleFunc("/auth", controller.Auth)
}
