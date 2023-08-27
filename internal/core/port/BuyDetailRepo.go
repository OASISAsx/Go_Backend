package port

type BuyDetailRepo interface {
	GetAll() ([]BuyDetail, error)
	Search(string) ([]BuyDetail, error)
	GetAllId(int) ([]BuyDetail, error)
	GetAllUserId(int) ([]BuyDetail, error)
	GetById(id int) (*BuyDetail, error)
	Create(BuyDetail) (*BuyDetail, error)
	Update(int, BuyDetail) error
	Delete(int) error
}

type BuyDetail struct {
	PayId         uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId        uint   `gorm:"notnull;type:int(10)"`
	ProductId     uint   `gorm:"notnull;type:int(10)"`
	ById          uint   `gorm:"notnull;type:int(10)"`
	PaySlip       string `gorm:"notnull"`
	ParNum        string
	TransPort     string 
	PayStatus     string `gorm:"notnull;default=กำลังดำเนินการ;type:varchar(20)"`
	ProductName   string `gorm:"notnull;type:varchar(50)"`
	ProductImages string
	ProductStock  string
	ProductPrice  int `gorm:"notnull;type:int(10)"`
	ProductType   string
	ProductDesc   string `gorm:"notnull;type:varchar(100)"`
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c BuyDetail) TableName() string {
	return "tbl_BuyDetail"
}
