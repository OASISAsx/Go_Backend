package port

type ProductRepo interface {
	GetAll() ([]Product, error)
	GetById(int) (*Product, error)
	Create(Product) (*Product, error)
	Update(int, Product) error
	UpdateSellStatus(int, bool) error
	Delete(int) error
	Search(string) ([]Product, error)
}

type Product struct {
	ProductId     uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	SvcId         uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	ProductName   string `gorm:"notnull;type:varchar(20)"`
	ProductDesc   string `gorm:"notnull;type:varchar(500)"`
	ProductType   string
	ProductPrice  int    `gorm:"type:int(20)"`
	ProductStock  string `gorm:"notnull;default=1;type:varchar(1)"`
	ProductImages string
	ProductImagex string
	ProductImagey string
	ProductImagez string
	RvRank        string `gorm:"notnull;type:varchar(5)"`
	Status        string `gorm:"notnull;default=กำลังดำเนินการ;type:varchar(20)"`
	SellStatus    bool   `gorm:"notnull"`
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c Product) TableName() string {
	return "tbl_product"
}
