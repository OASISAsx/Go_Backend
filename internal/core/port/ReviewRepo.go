package port

type ReviewRepo interface {
	GetAll() ([]Review, error)
	GetById(Review int) (*Review, error)
	Create(Review) (*Review, error)
	Update(int, Review) error
	Delete(int) error
}

type Review struct {
	RvId        uint `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId      uint
	ProductId   uint
	RvImg       string
	RvRank      string
	RvComment   string
	CreatedBy   string
	CreatedDate string
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Review) TableName() string {
	return "tbl_Review"
}
