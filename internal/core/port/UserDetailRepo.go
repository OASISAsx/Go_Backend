package port

type UserDetailRepo interface {
	GetAll() ([]UserDetail, error)
	GetById(id int) (*UserDetail, error)
	Create(UserDetail) (*UserDetail, error)
	Update(int, UserDetail) error
	Delete(int) error
}

type UserDetail struct {
	UserdtId        string `gorm:"primaryKey;autoIncrement;type:int(10)"`
	CardId          string
	SvcId           string
	FristNameUser   string
	LastNameUser    string
	FristNameSeller string
	LastNameSeller  string
	UserAddress     string
	UserZibId       string
	TCode           string
	ACode           string
	ProvinceId      string
	UserPhone       string
	UserEmail       string
	UserImege       string
	RecordStatus    string
	UserId          string
	CreatedBy       string
	CreatedDate     string
	UpdatedBy       string `gorm:"null"`
	UpdatedDate     string `gorm:"null"`
}

func (c UserDetail) TableName() string {
	return "tbl_UserDetail"
}
