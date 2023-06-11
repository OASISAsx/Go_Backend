package domain

type AreaapSvc interface {
	GetAllAreaap() ([]AreaapRespone, error)
	GetAreaap(int) (*AreaapRespone, error)
	AddAreaap(AreaapRequest) (*AreaapRespone, error)
	UpdateAreaap(int, AreaapRequest) error
	DeleteAreaap(int) error
}

type AreaapRequest struct {
	AmphurName  string `json:"Amphurname"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}

type AreaapRespone struct {
	AmphurId    uint   `json:"amphurid"`
	AmphurName  string `json:"Amphurname"`
	ProvinceId  string `json:"provinceId"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedDate string `json:"updatedDate"`
}
