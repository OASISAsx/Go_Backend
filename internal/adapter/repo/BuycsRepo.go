package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type buycsRepo struct {
	db *gorm.DB
}

func NewBuycsRepo(db *gorm.DB) port.BuycsRepo {
	return buycsRepo{
		db: db,
	}
}

func (c buycsRepo) GetAll() ([]port.Buycs, error) {
	buycss := []port.Buycs{}
	err := c.db.Find(&buycss).Error
	if err != nil {
		return nil, err
	}
	return buycss, nil
}

func (c buycsRepo) GetById(id int) (*port.Buycs, error) {
	buycs := port.Buycs{}
	err := c.db.First(&buycs, id).Error
	if err != nil {
		return nil, err
	}
	return &buycs, nil
}

func (c buycsRepo) Create(buycs port.Buycs) (*port.Buycs, error) {
	err := c.db.Create(&buycs).Error
	if err != nil {
		return nil, err
	}
	return &buycs, nil
}

func (c buycsRepo) Update(id int, buycs port.Buycs) error {
	err := c.db.Model(&port.Buycs{}).Where("rserve_id = ?", id).Updates(buycs).Error
	if err != nil {
		return err
	}
	return nil
}

func (c buycsRepo) Delete(id int) error {
	err := c.db.Delete(&port.Buycs{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
