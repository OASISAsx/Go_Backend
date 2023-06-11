package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type areapvRepo struct {
	db *gorm.DB
}

func NewAreapvRepo(db *gorm.DB) port.AreapvRepo {
	return areapvRepo{
		db: db,
	}
}

func (c areapvRepo) GetAll() ([]port.Areapv, error) {
	areapvs := []port.Areapv{}
	err := c.db.Find(&areapvs).Error
	if err != nil {
		return nil, err
	}
	return areapvs, nil
}

func (c areapvRepo) GetById(id int) (*port.Areapv, error) {
	areapv := port.Areapv{}
	err := c.db.First(&areapv, id).Error
	if err != nil {
		return nil, err
	}
	return &areapv, nil
}

func (c areapvRepo) Create(areapv port.Areapv) (*port.Areapv, error) {
	err := c.db.Create(&areapv).Error
	if err != nil {
		return nil, err
	}
	return &areapv, nil
}

func (c areapvRepo) Update(id int, areapv port.Areapv) error {
	err := c.db.Model(&port.Areapv{}).Where("province_id = ?", id).Updates(areapv).Error
	if err != nil {
		return err
	}
	return nil
}

func (c areapvRepo) Delete(id int) error {
	err := c.db.Delete(&port.Areapv{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
