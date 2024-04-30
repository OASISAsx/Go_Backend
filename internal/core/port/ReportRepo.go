package port

type ReportRepo interface {
	GetAll() ([]Report, error)
	GetById(id int) (*Report, error)
	Create(Report) (*Report, error)
	Update(int, Report) error
	Delete(int) error
}

type Report struct {
	ReportId uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	Username string `gorm:"notnull;type:varchar(50)"`
	Email    string `gorm:"notnull;type:varchar(30)"`
	Message  string `gorm:"notnull;type:varchar(100)"`
}

func (c Report) TableName() string {
	return "tbl_report"
}
