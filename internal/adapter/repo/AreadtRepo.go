package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type areadtRepo struct {
	db *gorm.DB
}

func NewAreadtRepo(db *gorm.DB) port.AreadtRepo {
	return areadtRepo{
		db: db,
	}
}

func (c areadtRepo) GetAll() ([]port.Areadt, error) {
	areadts := []port.Areadt{}
	err := c.db.Find(&areadts).Error
	if err != nil {
		return nil, err
	}
	return areadts, nil
}

func (c areadtRepo) GetById(id int) (*port.Areadt, error) {
	areadt := port.Areadt{}
	err := c.db.First(&areadt, id).Error
	if err != nil {
		return nil, err
	}
	return &areadt, nil
}

func (c areadtRepo) Create(areadt port.Areadt) (*port.Areadt, error) {
	err := c.db.Create(&areadt).Error
	if err != nil {
		return nil, err
	}
	return &areadt, nil
}

func (c areadtRepo) Update(id int, areadt port.Areadt) error {
	err := c.db.Model(&port.Areadt{}).Where("district_id = ?", id).Updates(areadt).Error
	if err != nil {
		return err
	}
	return nil
}

func (c areadtRepo) Delete(id int) error {
	err := c.db.Delete(&port.Areadt{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
