package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepo(db *gorm.DB) port.ReportRepo {
	return reportRepo{
		db: db,
	}
}

func (c reportRepo) GetAll() ([]port.Report, error) {
	reports := []port.Report{}
	err := c.db.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (c reportRepo) GetById(id int) (*port.Report, error) {
	report := port.Report{}
	err := c.db.First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (c reportRepo) Create(report port.Report) (*port.Report, error) {
	err := c.db.Create(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (c reportRepo) Update(id int, report port.Report) error {
	err := c.db.Model(&port.Report{}).Where("id = ?", id).Updates(report).Error
	if err != nil {
		return err
	}
	return nil
}

func (c reportRepo) Delete(id int) error {
	err := c.db.Delete(&port.Report{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
