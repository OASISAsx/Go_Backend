package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type buydetailRepo struct {
	db *gorm.DB
}

func NewBuyDetailRepo(db *gorm.DB) port.BuyDetailRepo {
	return buydetailRepo{
		db: db,
	}
}

func (c buydetailRepo) GetAll() ([]port.BuyDetail, error) {
	buydetails := []port.BuyDetail{}
	err := c.db.Find(&buydetails).Error
	if err != nil {
		return nil, err
	}
	return buydetails, nil
}

func (c buydetailRepo) GetById(id int) (*port.BuyDetail, error) {
	buydetail := port.BuyDetail{}
	err := c.db.First(&buydetail, id).Error
	if err != nil {
		return nil, err
	}
	return &buydetail, nil
}

func (c buydetailRepo) Create(buydetail port.BuyDetail) (*port.BuyDetail, error) {
	err := c.db.Create(&buydetail).Error
	if err != nil {
		return nil, err
	}
	return &buydetail, nil
}

func (c buydetailRepo) Update(id int, buydetail port.BuyDetail) error {
	err := c.db.Model(&port.BuyDetail{}).Where("buy_id = ?", id).Updates(buydetail).Error
	if err != nil {
		return err
	}
	return nil
}

func (c buydetailRepo) Delete(id int) error {
	err := c.db.Delete(&port.BuyDetail{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
