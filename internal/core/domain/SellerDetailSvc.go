package domain

type SellerDetailSvc interface {
	GetAllSellerDetail() ([]SellerDetailRespone, error)
	GetSellerDetail(int) (*SellerDetailRespone, error)
	AddSellerDetail(SellerDetailRequest) (*SellerDetailRespone, error)
	UpdateSellerDetail(int, SellerDetailRequest) error
	DeleteSellerDetail(int) error
}

type SellerDetailRequest struct {
	UserId       uint   `json:"userid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Phone        string `json:"phone"`
	BankName     string `json:"bankname"`
	BankId       string `json:"bankid"`
	PersonCard   string `json:"personcard"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdby"`
	CreatedDate  string `json:"createddate"`
	UpdatedBy    string `json:"updatedby"`
	UpdatedDate  string `json:"updateddate"`
}

type SellerDetailRespone struct {
	UserdeId     uint   `json:"userdeid"`
	UserId       uint   `json:"userid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Phone        string `json:"phone"`
	BankName     string `json:"bankname"`
	BankId       string `json:"bankid"`
	PersonCard   string `json:"personcard"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdby"`
	CreatedDate  string `json:"createddate"`
	UpdatedBy    string `json:"updatedby"`
	UpdatedDate  string `json:"updateddate"`
}
