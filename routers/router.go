package routers

import (
	"GoServer/controllers"
	"GoServer/service"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.CtrlGet("/", controllers.HomeController.Get)
	web.CtrlGet("/api/v1/account/", controllers.AccountController.GetAll)
	web.CtrlPost("/api/v1/account/register", controllers.AccountController.Register)
	web.CtrlPost("/api/v1/account/login", controllers.AccountController.Login)
	web.CtrlGet("/api/v1/account/logout", controllers.AccountController.Logout)
	web.CtrlGet("/ws", (*service.ClientController).ServeWs)

}
