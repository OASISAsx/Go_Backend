package httpserv

import (
	"collection/infrastructure"
	"collection/internal/adapter/handler"
	"collection/internal/adapter/repo"
	"collection/internal/core/service"
	"collection/middleware"

	"github.com/gin-gonic/gin"
)

func bindCustomer(g gin.Engine) {
	repo := repo.NewCustomerRepo(infrastructure.DB)
	svc := service.NewCustomerSvc(repo)
	hdl := handler.NewCustomerHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/customers", hdl.GetCustomers)
		v1.GET("/customers/:customerID", hdl.GetCustomer)
		v1.POST("/customers", hdl.AddCustomer)
		v1.PUT("/customers/:customerID", hdl.UpdateCustomer)
		v1.DELETE("/customers/:customerID", hdl.DeleteCustomer)
	}

}
func bindRole(g gin.Engine) {
	repo := repo.NewRoleRepo(infrastructure.DB)
	svc := service.NewRoleSvc(repo)
	hdl := handler.NewRoleHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/roles", hdl.GetRoles)
		v1.GET("/role/:roleID", hdl.GetRole)
		v1.POST("/role", hdl.AddRole)
		v1.PUT("/role/:roleID", hdl.UpdateRole)
		v1.DELETE("/role/:roleID", hdl.DeleteRole)
	}

}
func bindReview(g gin.Engine) {
	repo := repo.NewReviewRepo(infrastructure.DB)
	svc := service.NewReviewSvc(repo)
	hdl := handler.NewReviewHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/reviews", hdl.GetReviews)
		v1.GET("/review/:reviewID", hdl.GetReview)
		v1.POST("/review", hdl.AddReview)
		v1.PUT("/review/:reviewID", hdl.UpdateReview)
		v1.DELETE("/review/:reviewID", hdl.DeleteReview)
	}

}
func bindProduct(g gin.Engine) {
	repo := repo.NewProductRepo(infrastructure.DB)
	svc := service.NewProductSvc(repo)
	hdl := handler.NewProductHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/products", hdl.GetProducts)
		v1.GET("/product/:productID", hdl.GetProduct)
		v1.POST("/product", hdl.AddProduct)
		v1.PUT("/product/:productID", hdl.UpdateProduct)
		v1.DELETE("/product/:productID", hdl.DeleteProduct)
	}

}
func bindBuycs(g gin.Engine) {
	repo := repo.NewBuycsRepo(infrastructure.DB)
	svc := service.NewBuycsSvc(repo)
	hdl := handler.NewBuycsHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/buycss", hdl.GetBuycss)
		v1.GET("/buycs/:buycsID", hdl.GetBuycs)
		v1.POST("/buycs", hdl.AddBuycs)
		v1.PUT("/buycs/:buycsID", hdl.UpdateBuycs)
		v1.DELETE("/buycs/:buycsID", hdl.DeleteBuycs)
	}

}
func bindBuyDetail(g gin.Engine) {
	repo := repo.NewBuyDetailRepo(infrastructure.DB)
	svc := service.NewBuyDetailSvc(repo)
	hdl := handler.NewBuyDetailHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/buydetails", hdl.GetBuyDetails)
		v1.GET("/buydetail/:buydetailID", hdl.GetBuyDetail)
		v1.POST("/buydetail", hdl.AddBuyDetail)
		v1.PUT("/buydetail/:buydetailID", hdl.UpdateBuyDetail)
		v1.DELETE("/buydetail/:buydetailID", hdl.DeleteBuyDetail)
	}

}
func bindAreapv(g gin.Engine) {
	repo := repo.NewAreapvRepo(infrastructure.DB)
	svc := service.NewAreapvSvc(repo)
	hdl := handler.NewAreapvHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/areapvs", hdl.GetAreapvs)
		v1.GET("/areapv/:areapvID", hdl.GetAreapv)
		v1.POST("/areapv", hdl.AddAreapv)
		v1.PUT("/areapv/:areapvID", hdl.UpdateAreapv)
		v1.DELETE("/areapv/:areapvID", hdl.DeleteAreapv)
	}

}

func bindUserDetail(g gin.Engine) {
	repo := repo.NewUserDetailRepo(infrastructure.DB)
	svc := service.NewUserDetailSvc(repo)
	hdl := handler.NewUserDetailHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/userdetails", hdl.GetUserDetails)
		v1.GET("/userdetail/:userdetailID", hdl.GetUserDetail)
		v1.POST("/userdetail", hdl.AddUserDetail)
		v1.PUT("/userdetail/:userdetailID", hdl.UpdateUserDetail)
		v1.DELETE("/userdetail/:userdetailID", hdl.DeleteUserDetail)
	}

}
func bindAreaap(g gin.Engine) {
	repo := repo.NewAreaapRepo(infrastructure.DB)
	svc := service.NewAreaapSvc(repo)
	hdl := handler.NewAreaapHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/areaaps", hdl.GetAreaaps)
		v1.GET("/areaap/:areaapID", hdl.GetAreaap)
		v1.POST("/areaap", hdl.AddAreaap)
		v1.PUT("/areaap/:areaapID", hdl.UpdateAreaap)
		v1.DELETE("/areaap/:areaapID", hdl.DeleteAreaap)
	}
}
func bindAreadt(g gin.Engine) {
	repo := repo.NewAreadtRepo(infrastructure.DB)
	svc := service.NewAreadtSvc(repo)
	hdl := handler.NewAreadtHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/areadts", hdl.GetAreadts)
		v1.GET("/areadt/:areadtID", hdl.GetAreadt)
		v1.POST("/areadt", hdl.AddAreadt)
		v1.PUT("/areadt/:areadtID", hdl.UpdateAreadt)
		v1.DELETE("/areadt/:areadtID", hdl.DeleteAreadt)
	}
}
func bindRegister(g gin.Engine) {
	repo := repo.NewRegisterRepo(infrastructure.DB)
	svc := service.NewRegisterSvc(repo)
	hdl := handler.NewRegisterHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/registers", hdl.GetRegisters)
		v1.GET("/register/:registerID", hdl.GetRegister)
		v1.POST("/register", hdl.AddRegister)
		v1.PUT("/register/:registerID", hdl.UpdateRegister)
		v1.PUT("/user/role/:registerID", hdl.UpdateRole)
		v1.PUT("/user/status/:registerID", hdl.UpdateStatus)
		v1.DELETE("/register/:registerID", hdl.DeleteRegister)
		v1.POST("/login",hdl.Login)
		v1.GET("/user",middleware.DeserializeUser(repo),hdl.GetRegisters)
		v1.GET("/profile",middleware.DeserializeUser(repo),hdl.GetProfile)

	}
}

func bindImage(g gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.POST("/image", handler.FileUpload())
		v1.POST("/remote", handler.RemoteUpload())
	}
}
func bindCart(g gin.Engine) {
	repo := repo.NewCartRepo(infrastructure.DB)
	svc := service.NewCartSvc(repo)
	hdl := handler.NewCartHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/carts", hdl.GetCarts)
		v1.GET("/cart/:cartID", hdl.GetCart)
		v1.POST("/cart", hdl.AddCart)
		v1.PUT("/cart/:cartID", hdl.UpdateCart)
		v1.DELETE("/cart/:cartID", hdl.DeleteCart)
	}
}