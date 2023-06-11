package port

type AreadtRepo interface {
	GetAll() ([]Areadt, error)
	GetById(id int) (*Areadt, error)
	Create(Areadt) (*Areadt, error)
	Update(int, Areadt) error
	Delete(int) error
}

type Areadt struct {
	DistrictId   uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	DistrictCode string
	DistrictName string
	RefAmphurId  string
	CreatedBy    string
	CreatedDate  string
	UpdatedBy    string `gorm:"null"`
	UpdatedDate  string `gorm:"null"`
}

func (c Areadt) TableName() string {
	return "tbl_Areadt"
}
