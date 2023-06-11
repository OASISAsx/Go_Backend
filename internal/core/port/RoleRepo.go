package port

type RoleRepo interface {
	GetAll() ([]Role, error)
	GetById(id int) (*Role, error)
	Create(Role) (*Role, error)
	Update(int, Role) error
	Delete(int) error
}

type Role struct {
	RoleID      uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	RoleName    string
	RoleDesc    string
	Status      string
	CreatedBy   string
	CreatedDate string
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Role) TableName() string {
	return "tbl_role"
}
