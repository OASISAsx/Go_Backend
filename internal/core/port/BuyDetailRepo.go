package port

type BuyDetailRepo interface {
	GetAll() ([]BuyDetail, error)
	GetById(BuyDetail int) (*BuyDetail, error)
	Create(BuyDetail) (*BuyDetail, error)
	Update(int, BuyDetail) error
	Delete(int) error
}

type BuyDetail struct {
	BuyId      uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	StoreId    string
	ReserveId  string
	ProductId  string
	QrCode     string
	PayId      string
	CreatedBy   string
	CreatedDate string
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c BuyDetail) TableName() string {
	return "tbl_BuyDetail"
}
