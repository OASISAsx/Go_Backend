package repo

import (
	"collection/internal/core/port"
	"errors"

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
func (c productRepo) Search(productName string) ([]port.Product, error) {
	product := []port.Product{}
	result := c.db.Find(&product, "product_name LIKE ? OR product_type LIKE ? OR product_price LIKE ? OR product_desc LIKE ?", "%"+productName+"%", "%"+productName+"%", "%"+productName+"%", "%"+productName+"%")
	if result.Error != nil {
		return product, errors.New("product not found")
	}
	return product, nil
}

func (c productRepo) UpdateStatusProduct(id int, status bool) error {
	err := c.db.Model(&port.Product{}).Where("product_id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
func (c productRepo) UpdateSellStatus(id int, status bool) error {
	err := c.db.Model(&port.Product{}).Where("product_id = ?", id).Update("sell_status", status).Error
	if err != nil {
		return err
	}
	return nil
}
func (c productRepo) UpdateCount(id int, product port.Product) error {
	err := c.db.Model(&port.Product{}).Where("product_id = ?", id).Update("count", gorm.Expr("count + ?", product.Count)).Error
	if err != nil {
		return err
	}
	return nil
}