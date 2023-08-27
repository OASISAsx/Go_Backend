package domain

type UserDetailSvc interface {
	GetAllUserDetail() ([]UserDetailRespone, error)
	GetUserDetail(int) (*UserDetailRespone, error)
	AddUserDetail(UserDetailRequest) (*UserDetailRespone, error)
	UpdateUserDetail(int, UserDetailRequest) error
	DeleteUserDetail(int) error
}

type UserDetailRequest struct {
	CardId        string `json:"cardid"`
	UserId        uint   `json:"userid"`
	SvcId         string `json:"svcid"`
	FristNameUser string `json:"fristnameuser"`
	LastNameUser  string `json:"lastnameuser"`
	UserAddress   string `json:"useraddress"`
	UserZibId     string `json:"userzibId"`
	UserPhone     string `json:"userphone"`
	RecordStatus  string `json:"recordstatus"`
	Avatar        string `json:"Avatar"`
	Province      string `json:"province"`
	CreatedBy     string `json:"createdBy"`
	CreatedDate   string `json:"createdDate"`
	UpdatedBy     string `json:"updatedBy"`
	UpdatedDate   string `json:"updatedDate"`
}

type UserDetailRespone struct {
	UserdtId      uint   `json:"userdtid"`
	UserId        uint   `json:"userid"`
	CardId        string `json:"cardid"`
	SvcId         string `json:"svcid"`
	FristNameUser string `json:"fristnameuser"`
	LastNameUser  string `json:"lastnameuser"`
	UserAddress   string `json:"useraddress"`
	UserZibId     string `json:"userzibId"`
	UserPhone     string `json:"userphone"`
	Avatar        string `json:"Avatar"`
	Province      string `json:"province"`
	RecordStatus  string `json:"recordstatus"`
	CreatedBy     string `json:"createdBy"`
	CreatedDate   string `json:"createdDate"`
	UpdatedBy     string `json:"updatedBy"`
	UpdatedDate   string `json:"updatedDate"`
}
