package domain

type BuyDetailSvc interface {
	GetAllBuyDetail() ([]BuyDetailRespone, error)
	GetBuyDetail(int) (*BuyDetailRespone, error)
	AddBuyDetail(BuyDetailRequest) (*BuyDetailRespone, error)
	UpdateBuyDetail(int, BuyDetailRequest) error
	DeleteBuyDetail(int) error
}

type BuyDetailRequest struct {
	StoreId     string `json:"storeid"`
	ReserveId   string `json:"reserveid"`
	ProductId   string `json:"productid"`
	QrCode      string `json:"qrcode"`
	PayId       string `json:"payid"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type BuyDetailRespone struct {
	BuyId       uint   `json:"buyid"`
	StoreId     string `json:"storeid"`
	ReserveId   string `json:"reserveid"`
	ProductId   string `json:"productid"`
	QrCode      string `json:"qrcode"`
	PayId       string `json:"payid"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}
