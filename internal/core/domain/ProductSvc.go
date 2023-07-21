package domain

type ProductSvc interface {
	GetAllProduct() ([]ProductRespone, error)
	GetById(int) (*ProductRespone, error)
	AddProduct(ProductRequest) (*ProductRespone, error)
	UpdateProduct(int, ProductRequest) error
	DeleteProduct(int) error
	Search(string) (*[]ProductRespone, error)
	UpdateStatusProduct(int, StatusProduct) error
	UpdateSellStatus(int, SellStatusProduct) error
	UpdateCount(int, ProductRequest) error
}

type ProductRequest struct {
	SvcId         uint   `json:"svcid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductPrice  string `json:"productprice"`
	ProductStock  string `json:"productstock"`
	ProductImages string `json:"productimages"`
	Producttype   string `json:"producttype"`
	Status        bool   `json:"status"`
	SellStatus    bool   `json:"sellstatus"`
	Count         int    `json:"count"`
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
	ProductPrice  string `json:"productprice"`
	ProductStock  string `json:"productstock"`
	ProductImages string `json:"productimages"`
	Producttype   string `json:"producttype"`
	Status        bool   `json:"status"`
	SellStatus    bool   `json:"sellstatus"`
	Count         int    `json:"count"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
	ProductQr     string `json:"productqr"`
}
type StatusProduct struct {
	Status bool `json:"status"`
}
type SellStatusProduct struct {
	SellStatus bool `json:"sellstatus"`
}
