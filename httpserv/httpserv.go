package httpserv

import (
	"collection/app"
	"collection/middleware"
	"fmt"

	"github.com/spf13/viper"
)

func Run() {
	a := app.NewApp()
	m := middleware.New()
	a.UseMiddleware(m.Logger)
	a.UseMiddleware(m.ErrorHandler)
	a.UseMiddleware(m.Cors)
	bindCustomer(a.GinEngine())
	bindRole(a.GinEngine())
	bindReview(a.GinEngine())
	bindProduct(a.GinEngine())
	bindBuycs(a.GinEngine())
	bindBuyDetail(a.GinEngine())
	bindAreapv(a.GinEngine())
	bindUserDetail(a.GinEngine())
	bindAreaap(a.GinEngine())
	bindAreadt(a.GinEngine())
	bindRegister(a.GinEngine())
	bindImage(a.GinEngine())
	bindSellerDetail(a.GinEngine())

	bindEmail(a.GinEngine())

	port := fmt.Sprintf(":%v", viper.GetInt("app.port"))
	a.Start(port)
}
