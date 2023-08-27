package domain

type BuyDetailSvc interface {
	GetAllBuyDetail() ([]BuyDetailRespone, error)
	SearchBuyDetail(string) (*[]BuyDetailRespone, error)
	GetAllBuyDetailId(int) ([]BuyDetailRespone, error)
	GetAllUserId(int) ([]BuyDetailRespone, error)
	GetBuyDetail(int) (*BuyDetailRespone, error)
	AddBuyDetail(BuyDetailRequest) (*BuyDetailRespone, error)
	UpdateBuyDetail(int, BuyDetailRequest) error
	DeleteBuyDetail(int) error
}

type BuyDetailRequest struct {
	UserId        uint   `json:"userid"`
	ProductId     uint   `json:"productid"`
	ById          uint   `json:"byid"`
	PaySlip       string `json:"payslip"`
	ParNum        string `json:"parnum"`
	TransPort     string `json:"transport"`
	PayStatus     string `json:"paymentstatus"`
	ProductName   string `json:"productname"`
	ProductImages string `json:"productimages"`
	ProductPrice  int    `json:"productprice"`
	ProductStock  string    `json:"productstock"`
	ProductType   string `json:"producttype"`
	ProductDesc   string `json:"productdesc"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}

type BuyDetailRespone struct {
	PayId         uint   `json:"payid"`
	UserId        uint   `json:"userid"`
	ProductId     uint   `json:"productid"`
	ById          uint   `json:"byid"`
	PaySlip       string `json:"payslip"`
	ParNum        string `json:"parnum"`
	TransPort     string `json:"transport"`
	PayStatus     string `json:"paymentstatus"`
	ProductName   string `json:"productname"`
	ProductImages string `json:"productimages"`
	ProductPrice  int    `json:"productprice"`
	ProductStock  string    `json:"productstock"`
	ProductType   string `json:"producttype"`
	ProductDesc   string `json:"productdesc"`
	CreatedBy     string `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	UpdatedBy     string `json:"updatedby"`
	UpdatedDate   string `json:"updateddate"`
}
