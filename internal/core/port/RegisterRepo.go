package port

type RegisterRepo interface {
	FindByUsername(username string) (Register, error)
	GetAll() ([]Register, error)
	GetById(id int) (*Register, error)
	Create(Register) (*Register, error)
	Update(int, Register) error
	Delete(int) error
}

type Register struct {
	UserId       uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	RoleId       int
	Username     string
	Password     string
	Nickname     string
	Email        string
	RecordStatus string
	CreatedBy    string
	CreatedDate  string
	UpdatedBy    string `gorm:"null"`
	UpdatedDate  string `gorm:"null"`
}

func (u Register) TableName() string {
	return "tbl_User"
}
