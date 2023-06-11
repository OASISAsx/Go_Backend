package domain

type UserDetailSvc interface {
	GetAllUserDetail() ([]UserDetailRespone, error)
	GetUserDetail(int) (*UserDetailRespone, error)
	AddUserDetail(UserDetailRequest) (*UserDetailRespone, error)
	UpdateUserDetail(int, UserDetailRequest) error
	DeleteUserDetail(int) error
}

type UserDetailRequest struct {
	CardId          string `json:"cardid"`
	SvcId           string `json:"svcid"`
	FristNameUser   string `json:"fristnameuser"`
	LastNameUser    string `json:"lastnameuser"`
	FristNameSeller string `json:"fristnameseller"`
	LastNameSeller  string `json:"lastnameseller"`
	UserAddress     string `json:"useraddress"`
	UserZibId       string `json:"userzibId"`
	TCode           string `json:"tcode"`
	ACode           string `json:"acode"`
	ProvinceId      string `json:"provinceid"`
	UserPhone       string `json:"userphone"`
	UserEmail       string `json:"useremail"`
	UserImege       string `json:"userimege"`
	RecordStatus    string `json:"recordstatus"`
	UserId          string `json:"userid"`
	CreatedBy       string `json:"createdBy"`
	CreatedDate     string `json:"createdDate"`
	UpdatedBy       string `json:"updatedBy"`
	UpdatedDate     string `json:"updatedDate"`
}

type UserDetailRespone struct {
	UserdtId        string `json:"userdtid"`
	CardId          string `json:"cardid"`
	SvcId           string `json:"svcid"`
	FristNameUser   string `json:"fristnameuser"`
	LastNameUser    string `json:"lastnameuser"`
	FristNameSeller string `json:"fristnameseller"`
	LastNameSeller  string `json:"lastnameseller"`
	UserAddress     string `json:"useraddress"`
	UserZibId       string `json:"userzibId"`
	TCode           string `json:"tcode"`
	ACode           string `json:"acode"`
	ProvinceId      string `json:"provinceid"`
	UserPhone       string `json:"userphone"`
	UserEmail       string `json:"useremail"`
	UserImege       string `json:"userimege"`
	RecordStatus    string `json:"recordstatus"`
	UserId          string `json:"userid"`
	CreatedBy       string `json:"createdBy"`
	CreatedDate     string `json:"createdDate"`
	UpdatedBy       string `json:"updatedBy"`
	UpdatedDate     string `json:"updatedDate"`
}
