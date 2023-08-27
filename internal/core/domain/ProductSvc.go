package domain

type ProductSvc interface {
	GetAllProduct() ([]ProductRespone, error)
	GetById(int) (*ProductRespone, error)
	AddProduct(ProductRequest) (*ProductRespone, error)
	UpdateProduct(int, ProductRequest) error
	UpdateSellStatus(int, SellStatusProduct) error
	DeleteProduct(int) error
	Search(string) (*[]ProductRespone, error)
}

type ProductRequest struct {
	SvcId         uint   `json:"svcid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductPrice  int    `json:"productprice"`
	ProductStock  string `json:"productstock"`
	ProductImages string `json:"productimages"`
	ProductImagex string `json:"productimagex"`
	ProductImagey string `json:"productimagey"`
	ProductImagez string `json:"productimagez"`
	ProductType   string `json:"producttype"`
	Status        string `json:"status"`
	SellStatus    bool   `json:"sellstatus"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}

type ProductRespone struct {
	ProductId     uint   `json:"productid"`
	SvcId         uint   `json:"svcid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductPrice  int    `json:"productprice"`
	ProductStock  string `json:"productstock"`
	ProductImages string `json:"productimages"`
	ProductImagex string `json:"productimagex"`
	ProductImagey string `json:"productimagey"`
	ProductImagez string `json:"productimagez"`
	ProductType   string `json:"producttype"`
	Status        string `json:"status"`
	SellStatus    bool   `json:"sellstatus"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
	ProductQr     string `json:"productqr"`
}
type SellStatusProduct struct {
	SellStatus bool `json:"sellstatus"`
}