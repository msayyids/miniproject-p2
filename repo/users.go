package repo

import (
	"fmt"
	"miniproject/model"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r *Repo) AddUsers(name, email, password, phoneNumber string) (model.Users, error) {

	inputUser := model.Users{Name: name, Email: email, Password: password, PhoneNumber: phoneNumber}
	query := r.DB.Create(&inputUser)
	if query.Error != nil {
		return model.Users{}, query.Error
	}

	return inputUser, nil
}

func (r *Repo) FindUserByEmail(email string) (model.Users, error) {
	user := model.Users{}
	query := r.DB.First(&user, "email=?", email)
	if query.Error != nil {
		return model.Users{}, query.Error
	}

	return user, nil
}

func (r *Repo) FindById(id int) (model.Users, error) {
	user := model.Users{}
	query := r.DB.First(&user, "id=?", id)
	if query.Error != nil {
		return model.Users{}, query.Error
	}

	return user, nil
}

func (r *Repo) EditAmount(amount int, userid int) error {
	var user model.Users
	if err := r.DB.Where("id = ?", userid).First(&user).Error; err != nil {
		return fmt.Errorf("user not found")
	}

	user.Deposit_amount += amount

	if err := r.DB.Model(&user).Update("deposit_amount", user.Deposit_amount).Error; err != nil {
		return err
	}

	return nil
}
