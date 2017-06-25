package routers

import (
	"github.com/astaxie/beego"
	"github.com/ebitgo/quotes/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getprice", &controllers.GetPriceController{})
}
