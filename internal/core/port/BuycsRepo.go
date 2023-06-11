package port

type BuycsRepo interface {
	GetAll() ([]Buycs, error)
	GetById(id int) (*Buycs, error)
	Create(Buycs) (*Buycs, error)
	Update(int, Buycs) error
	Delete(int) error
}

type Buycs struct {
	RserveId     uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	StoreId      string
	UserId       string
	RecordStatus string
	CreatedBy     string
	CreatedDate   string
	UpdatedBy     string `gorm:"null"`
	UpdatedDate   string `gorm:"null"`
}

func (c Buycs) TableName() string {
	return "tbl_Buycs"
}
