package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type sellerdetailRepo struct {
	db *gorm.DB
}

func NewSellerDetailRepo(db *gorm.DB) port.SellerDetailRepo {
	return sellerdetailRepo{
		db: db,
	}
}

func (c sellerdetailRepo) GetAll() ([]port.SellerDetail, error) {
	sellerdetails := []port.SellerDetail{}
	err := c.db.Find(&sellerdetails).Error
	if err != nil {
		return nil, err
	}
	return sellerdetails, nil
}

func (c sellerdetailRepo) GetById(id int) (*port.SellerDetail, error) {
	sellerdetail := port.SellerDetail{}
	err := c.db.Where("user_id = ?", id).First(&sellerdetail).Error
	if err != nil {
		return nil, err
	}
	return &sellerdetail, nil
}

func (c sellerdetailRepo) Create(sellerdetail port.SellerDetail) (*port.SellerDetail, error) {
	err := c.db.Create(&sellerdetail).Error
	if err != nil {
		return nil, err
	}
	return &sellerdetail, nil
}

func (c sellerdetailRepo) Update(id int, sellerdetail port.SellerDetail) error {
	err := c.db.Model(&port.SellerDetail{}).Where("user_id = ?", id).Updates(sellerdetail).Error
	if err != nil {
		return err
	}
	return nil
}

func (c sellerdetailRepo) Delete(id int) error {
	err := c.db.Delete(&port.SellerDetail{}, id).Error
	if err != nil {
		return err
	}
	return nil
}