package domain

type CartSvc interface {
	GetAllCart() ([]CartRespone, error)
	GetById(int) ([]CartRespone, error)
	AddCart(CartRequest) (*CartRespone, error)
	UpdateCart(int, CartRequest) error
	DeleteCart(int) error
}

type CartRequest struct {
	SvcId         uint   `json:"svcid"`
	UserId        uint   `json:"userid"`
	ProductId     uint   `json:"productid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductImages string `json:"productimages"`
	ProductType   string `json:"producttype"`
	ProductPrice  int    `json:"productpice"`
	ProductStock  string `json:"productstock"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}

type CartRespone struct {
	SvcId         uint   `json:"svcid"`
	CartId        uint   `json:"cartid"`
	UserId        uint   `json:"userid"`
	ProductId     uint   `json:"productid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductImages string `json:"productimages"`
	ProductType   string `json:"producttype"`
	ProductPrice  int    `json:"productprice"`
	ProductStock  string `json:"productstock"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}
