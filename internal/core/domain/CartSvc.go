package domain

type CartSvc interface {
	GetAllCart() ([]CartRespone, error)
	GetById(int) (*CartRespone, error)
	AddCart(CartRequest) (*CartRespone, error)
	UpdateCart(int, CartRequest) error
	DeleteCart(int) error
}

type CartRequest struct {
	SvcId         uint   `json:"svcid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductImages string `json:"productimages"`
	Producttype   string `json:"producttype"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}

type CartRespone struct {
	CartId        uint   `json:"cartid"`
	SvcId         uint   `json:"svcid"`
	ProductName   string `json:"productname"`
	ProductDesc   string `json:"productdesc"`
	ProductImages string `json:"productimages"`
	Producttype   string `json:"producttype"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}
