package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type areaapRepo struct {
	db *gorm.DB
}

func NewAreaapRepo(db *gorm.DB) port.AreaapRepo {
	return areaapRepo{
		db: db,
	}
}

func (c areaapRepo) GetAll() ([]port.Areaap, error) {
	areaaps := []port.Areaap{}
	err := c.db.Find(&areaaps).Error
	if err != nil {
		return nil, err
	}
	return areaaps, nil
}

func (c areaapRepo) GetById(id int) (*port.Areaap, error) {
	areaap := port.Areaap{}
	err := c.db.First(&areaap, id).Error
	if err != nil {
		return nil, err
	}
	return &areaap, nil
}

func (c areaapRepo) Create(areaap port.Areaap) (*port.Areaap, error) {
	err := c.db.Create(&areaap).Error
	if err != nil {
		return nil, err
	}
	return &areaap, nil
}

func (c areaapRepo) Update(id int, areaap port.Areaap) error {
	err := c.db.Model(&port.Areaap{}).Where("amphur_id = ?", id).Updates(areaap).Error
	if err != nil {
		return err
	}
	return nil
}

func (c areaapRepo) Delete(id int) error {
	err := c.db.Delete(&port.Areaap{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
