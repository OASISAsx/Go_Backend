package port

type UserDetailRepo interface {
	GetAll() ([]UserDetail, error)
	GetById(id int) (*UserDetail, error)
	Create(UserDetail) (*UserDetail, error)
	Update(int, UserDetail) error
	Delete(int) error
}

type UserDetail struct {
	UserdtId      uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId        uint
	CardId        string
	SvcId         string
	FristNameUser string
	LastNameUser  string
	UserAddress   string
	UserZibId     string
	Province      string
	UserPhone     string
	Avatar        string
	RecordStatus  string
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c UserDetail) TableName() string {
	return "tbl_UserDetail"
}
