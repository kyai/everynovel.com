package routers

import (
	"everynovel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/index", &controllers.IndexController{})
    beego.Router("/GetWord", &controllers.WordController{}, "get:GetWord")
    beego.Router("/PostWord", &controllers.WordController{}, "post:PostWord")
    beego.Router("/GetUser", &controllers.UserController{}, "get:GetUser")
    // beego.Router("/PostUser", &controllers.UserController{}, "post:PostUser")
}
