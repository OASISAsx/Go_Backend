package domain

type BuycsSvc interface {
	GetAllBuycs() ([]BuycsRespone, error)
	GetBuycs(int) (*BuycsRespone, error)
	AddBuycs(BuycsRequest) (*BuycsRespone, error)
	UpdateBuycs(int, BuycsRequest) error
	DeleteBuycs(int) error
}

type BuycsRequest struct {
	StoreId      string `json:"storeid"`
	UserId       string `json:"userid"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}

type BuycsRespone struct {
	RserveId     uint   `json:"rserveid"`
	StoreId      string `json:"storeid"`
	UserId       string `json:"userid"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}
