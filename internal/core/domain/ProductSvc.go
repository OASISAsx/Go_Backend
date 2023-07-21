package domain

type ProductSvc interface {
	GetAllProduct() ([]ProductRespone, error)
	GetById(int) (*ProductRespone, error)
	AddProduct(ProductRequest) (*ProductRespone, error)
	UpdateProduct(int, ProductRequest) error
	DeleteProduct(int) error
}

type ProductRequest struct {
	SvcId         uint   `json:"svcid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductPrice  uint   `json:"productprice"`
	ProductStock  uint   `json:"productstock"`
	ProductImages string `json:"productimages"`
	Producttype   string `json:"producttype"`
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
	ProductPrice  uint   `json:"productprice"`
	ProductStock  uint   `json:"productstock"`
	ProductImages string `json:"productimages"`
	Producttype   string `json:"producttype"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}
