package port

type UserDetailRepo interface {
	GetAll() ([]UserDetail, error)
	GetById(id int) (*UserDetail, error)
	Create(UserDetail) (*UserDetail, error)
	Update(int, UserDetail) error
	Delete(int) error
}

type UserDetail struct {
	UserdtId        uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId          uint
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
	BankName        string `gorm:"notnull;type:varchar(100)"`
	BankId          string `gorm:"notnull;type:varchar(10)"`
	PersonCard      string `gorm:"null;type:longtext"`
	RecordStatus    string
	CreatedBy       string
	CreatedDate     string
	UpdatedBy       string `gorm:"null"`
	UpdatedDate     string `gorm:"null"`
}

func (c UserDetail) TableName() string {
	return "tbl_UserDetail"
}
