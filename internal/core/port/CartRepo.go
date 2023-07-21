package port

type CartRepo interface {
	GetAll() ([]Cart, error)
	GetById(int) (*Cart, error)
	Create(Cart) (*Cart, error)
	Update(int, Cart) error
	Delete(int) error
}

type Cart struct {
	CartId        uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	SvcId         uint
	ProductName   string
	ProductDesc   string
	Producttype   string
	ProductImages string
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c Cart) TableName() string {
	return "tbl_cart"
}
