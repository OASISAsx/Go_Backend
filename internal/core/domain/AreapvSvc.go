package domain

type AreapvSvc interface {
	GetAllAreapv() ([]AreapvRespone, error)
	GetAreapv(int) (*AreapvRespone, error)
	AddAreapv(AreapvRequest) (*AreapvRespone, error)
	UpdateAreapv(int, AreapvRequest) error
	DeleteAreapv(int) error
}

type AreapvRequest struct {
	ProvinceName string `json:"provincename"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}

type AreapvRespone struct {
	ProvinceId   uint   `json:"provinceid"`
	ProvinceName string `json:"provincename"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}
