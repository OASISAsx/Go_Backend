package port

type ProductRepo interface {
	GetAll() ([]Product, error)
	GetById(int) (*Product, error)
	Create(Product) (*Product, error)
	Update(int, Product) error
	Delete(int) error
}

type Product struct {
	ProductId     uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	SvcId         string
	ProductName   string
	ProductDesc   string
	ProductImages string
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c Product) TableName() string {
	return "tbl_product"
}
