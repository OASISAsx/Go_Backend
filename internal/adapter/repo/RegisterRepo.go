package repo

import (
	"collection/internal/core/port"
	"errors"

	"gorm.io/gorm"
)

type registerRepo struct {
	db *gorm.DB
}

func NewRegisterRepo(db *gorm.DB) port.RegisterRepo {
	return registerRepo{
		db: db,
	}
}

func (c registerRepo) GetAll() ([]port.Register, error) {
	registers := []port.Register{}
	err := c.db.Find(&registers).Error
	if err != nil {
		return nil, err
	}
	return registers, nil
}

func (c registerRepo) GetById(id int) (*port.Register, error) {
	register := port.Register{}
	err := c.db.First(&register, id).Error
	if err != nil {
		return nil, err
	}
	return &register, nil
}
func (c registerRepo) Create(userExist port.Register) (*port.Register, error) {
	_ = c.db.Where("username = ?", userExist.Username).First(&userExist).Error
	if userExist.UserId > 0 {
		err := errors.New("username already exist")
		return nil, err
	} else {
		err := c.db.Create(&userExist).Error
		if err != nil {
			return &userExist, err
		}
	}
	return &userExist, nil
}

func (c registerRepo) Update(id int, register port.Register) error {
	err := c.db.Model(&port.Register{}).Where("user_id = ?", id).Updates(register).Error
	if err != nil {
		return err
	}
	return nil
}

func (c registerRepo) Delete(id int) error {
	err := c.db.Delete(&port.Register{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (c registerRepo) FindByUsername(username string) (port.Register, error) {
	var user port.Register
	result := c.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return user, errors.New("invalid username or Password")
	}
	return user, nil

}
