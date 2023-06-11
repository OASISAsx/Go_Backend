package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) port.ProductRepo {
	return productRepo{
		db: db,
	}
}

func (c productRepo) GetAll() ([]port.Product, error) {
	products := []port.Product{}
	err := c.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c productRepo) GetById(id int) (*port.Product, error) {
	product := port.Product{}
	err := c.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (c productRepo) Create(product port.Product) (*port.Product, error) {
	err := c.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (c productRepo) Update(id int, product port.Product) error {
	err := c.db.Model(&port.Product{}).Where("product_id = ?", id).Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (c productRepo) Delete(id int) error {
	err := c.db.Delete(&port.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
