package port

type ProductRepo interface {
	GetAll() ([]Product, error)
	GetById(int) (*Product, error)
	Create(Product) (*Product, error)
	Update(int, Product) error
	Delete(int) error
	Search(string) ([]Product, error)
	UpdateStatusProduct(int, bool) error
	UpdateCount(int, Product) error
	UpdateSellStatus(int, bool) error
}

type Product struct {
	ProductId     uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	SvcId         uint
	ProductName   string
	ProductDesc   string `gorm:"notnull;type:varchar(500)"`
	ProductType   string
	ProductPrice  int `gorm:"type:int(20)"`
	ProductStock  string
	ProductImages string
	Status        bool `gorm:"column:status;notnull"`
	SellStatus    bool `gorm:"notnull"`
	Count         int  `gorm:"type:int(10)"`
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c Product) TableName() string {
	return "tbl_product"
}
