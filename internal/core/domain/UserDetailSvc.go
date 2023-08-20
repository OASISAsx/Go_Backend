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
	UserId          uint   `json:"userid"`
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
	CreatedBy       string `json:"createdBy"`
	CreatedDate     string `json:"createdDate"`
	UpdatedBy       string `json:"updatedBy"`
	UpdatedDate     string `json:"updatedDate"`
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
	ProvinceId    string `json:"provinceid"`
	UserPhone     string `json:"userphone"`
	UserEmail     string `json:"useremail"`
	BankName      string `gorm:"notnull;type:varchar(100)"`
	BankId        string `gorm:"notnull;type:varchar(10)"`
	PersonCard    string `gorm:"null;type:longtext"`
	RecordStatus  string `json:"recordstatus"`
	CreatedBy     string `json:"createdBy"`
	CreatedDate   string `json:"createdDate"`
	UpdatedBy     string `json:"updatedBy"`
	UpdatedDate   string `json:"updatedDate"`
}
