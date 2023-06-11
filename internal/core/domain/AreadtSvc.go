package domain

type AreadtSvc interface {
	GetAllAreadt() ([]AreadtRespone, error)
	GetAreadt(int) (*AreadtRespone, error)
	AddAreadt(AreadtRequest) (*AreadtRespone, error)
	UpdateAreadt(int, AreadtRequest) error
	DeleteAreadt(int) error
}

type AreadtRequest struct {
	DistrictName string `json:"districtname"`
	RefAmphurId  string `json:"refamphurid"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}

type AreadtRespone struct {
	DistrictId   uint   `json:"districtid"`
	DistrictCode string `json:"districtcode"`
	DistrictName string `json:"districtname"`
	RefAmphurId  string `json:"refamphurid"`
	CreatedBy    string `json:"createdBy"`
	CreatedDate  string `json:"createdDate"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedDate  string `json:"updatedDate"`
}
