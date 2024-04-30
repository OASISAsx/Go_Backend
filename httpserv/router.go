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
		v1.PUT("review/status/:reviewID", hdl.UpdateReviewStatus)
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
		v1.PUT("/product/status/seller/:productID", hdl.UpdateSellStatus)
		v1.DELETE("/product/:productID", hdl.DeleteProduct)
		v1.GET("/search/:ProductName", hdl.Search)
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
		v1.GET("/buydetails/:buydetailID", hdl.GetAllBuyDetailId)
		v1.GET("/buydetails/search/:name", hdl.SearchBuyDetail)
		v1.GET("/buydetails/user/:UserId", hdl.GetAllUserId)
		v1.POST("/buydetail", hdl.AddBuyDetail)
		v1.PUT("/buydetail/:buydetailID", hdl.UpdateBuyDetail)
		v1.DELETE("/buydetail/:buydetailID", hdl.DeleteBuyDetail)
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
func bindSellerDetail(g gin.Engine) {
	repo := repo.NewSellerDetailRepo(infrastructure.DB)
	svc := service.NewSellerDetailSvc(repo)
	hdl := handler.NewSellerDetailHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/sellerdetails", hdl.GetSellerDetails)
		v1.GET("/sellerdetail/:sellerdetailID", hdl.GetSellerDetail)
		v1.POST("/sellerdetail", hdl.AddSellerDetail)
		v1.PUT("/sellerdetail/:sellerdetailID", hdl.UpdateSellerDetail)
		v1.DELETE("/sellerdetail/:sellerdetailID", hdl.DeleteSellerDetail)
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
		v1.PUT("/register/role/:registerID", hdl.UpdateRole)
		v1.PUT("/register/status/:registerID", hdl.UpdateStatus)
		v1.DELETE("/register/:registerID", hdl.DeleteRegister)
		v1.POST("/login", hdl.Login)
		v1.GET("/register", middleware.DeserializeUser(repo), hdl.GetRegisters)
		v1.GET("/profile", middleware.DeserializeUser(repo), hdl.GetProfile)

	}
}

func bindImage(g gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.POST("/image", handler.FileUpload())
		v1.POST("/remote", handler.RemoteUpload())

	}
}
func bindImagex(g gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.POST("/imagex", handler.FileUploads())
		v1.POST("/remotex", handler.RemoteUploads())

	}
}
func bindImagey(g gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.POST("/imagey", handler.FileUploady())
		v1.POST("/remotey", handler.RemoteUploady())

	}
}

func bindImagez(g gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.POST("/imagez", handler.FileUploadz())
		v1.POST("/remotez", handler.RemoteUploadz())

	}
}

func bindEmail(g gin.Engine) {
	svc := service.NewSenderSvc()
	hdl := handler.NewSenderHdl(svc)
	v1 := g.Group("/v1")
	{
		v1.POST("/sender", hdl.SendMail())
	}
}
func bindReprot(g gin.Engine) {
	repo := repo.NewReportRepo(infrastructure.DB)
	svc := service.NewReportSvc(repo)
	hdl := handler.NewReportHdl(svc)

	v1 := g.Group("/v1")
	{
		v1.GET("/reports", hdl.GetReports)
		v1.GET("/report/:reportID", hdl.GetReport)
		v1.POST("/report", hdl.AddReport)
		v1.PUT("/report/:reportID", hdl.UpdateReport)
		v1.DELETE("/report/:reportID", hdl.DeleteReport)
	}

}
