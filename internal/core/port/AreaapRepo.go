package port

type AreaapRepo interface {
	GetAll() ([]Areaap, error)
	GetById(id int) (*Areaap, error)
	Create(Areaap) (*Areaap, error)
	Update(int, Areaap) error
	Delete(int) error
}

type Areaap struct {
	AmphurId    uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	AmphurName  string
	ProvinceId  string
	CreatedBy   string
	CreatedDate string
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Areaap) TableName() string {
	return "tbl_Areaap"
}
