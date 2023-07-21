package repo

import (
	"collection/internal/core/port"

	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) port.CartRepo {
	return cartRepo{
		db: db,
	}
}

func (c cartRepo) GetAll() ([]port.Cart, error) {
	carts := []port.Cart{}
	err := c.db.Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c cartRepo) GetById(id int) ([]port.Cart, error) {
	cart := []port.Cart{}
	err := c.db.Where("user_id = ?", id).Find(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c cartRepo) Create(cart port.Cart) (*port.Cart, error) {
	err := c.db.Create(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c cartRepo) Update(id int, cart port.Cart) error {
	err := c.db.Model(&port.Cart{}).Where("cart_id = ?", id).Updates(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (c cartRepo) Delete(id int) error {
	err := c.db.Delete(&port.Cart{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
