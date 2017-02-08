package routers

import (
	"BW/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("home", &controllers.MainController{})
    beego.Router("citizen",&controllers.MainController{},"get:Citizen")
    beego.Router("hobby",&controllers.MainController{},"get:Hobby")
    beego.Router("moment",&controllers.MainController{},"get:Moment")
    beego.Router("sign",&controllers.MainController{},"get:Sign")
beego.Router("control_co",&controllers.Control_pannel{},"get:Control_co")
beego.Router("control_ft",&controllers.Control_pannel{},"get:Control_ft") 
beego.Router("control_hy",&controllers.Control_pannel{},"get:Control_hy") 
beego.Router("control_sx",&controllers.Control_pannel{},"get:Control_sx")
beego.Router("control_tj",&controllers.Control_pannel{},"get:Control_tj")
}
