package port

type AreapvRepo interface {
	GetAll() ([]Areapv, error)
	GetById(id int) (*Areapv, error)
	Create(Areapv) (*Areapv, error)
	Update(int, Areapv) error
	Delete(int) error
}

type Areapv struct {
	ProvinceId   uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	ProvinceName string
	CreatedBy    string
	CreatedDate  string
	UpdatedBy    string `gorm:"null"`
	UpdatedDate  string `gorm:"null"`
}

func (c Areapv) TableName() string {
	return "tbl_Areapv"
}
