package port

type SellerDetailRepo interface {
	GetAll() ([]SellerDetail, error)
	GetById(id int) (*SellerDetail, error)
	Create(SellerDetail) (*SellerDetail, error)
	Update(int, SellerDetail) error
	Delete(int) error
}

type SellerDetail struct {
	UserdeId     uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId       uint   `gorm:"notnull;type:int(10)"`
	FirstName    string `gorm:"notnull;type:varchar(100)"`
	LastName     string `gorm:"notnull;type:varchar(100)"`
	Phone        string `gorm:"notnull;type:varchar(10)"`
	BankName     string `gorm:"notnull;type:varchar(100)"`
	BankId       string `gorm:"notnull;type:varchar(10)"`
	PersonCard   string `gorm:"null;type:longtext"`
	RecordStatus string `gorm:"notnull;type:varchar(30)"`
	CreatedBy    string `gorm:"notnull;type:varchar(10)"`
	CreatedDate  string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy    string `gorm:"null;type:varchar(10)"`
	UpdatedDate  string `gorm:"null;type:varchar(20)"`
}

func (c SellerDetail) TableName() string {
	return "tbl_sellerdetails"
}