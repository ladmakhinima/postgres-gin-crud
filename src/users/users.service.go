package users

import (
	"fmt"

	"github.com/ladmakhinima/postgres-gin-crud/src/config"
	"golang.org/x/crypto/bcrypt"
)

func (user *User) SaveService() (*User, string) {
	var duplicateEmail User

	config.DB.Where("email = ?", user.Email).First(&duplicateEmail)

	if duplicateEmail.ID != 0 {
		fmt.Println("Duplicate Error User Email Address")
		return &User{}, "Duplicate Error User Email Address"
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	user.Password = string(password)

	config.DB.Create(user)

	return user, ""
}

func FindAllService() []User {
	var users []User
	config.DB.Order("id desc").Find(&users)
	return users
}

func FindByIDService(id int) (User, error) {
	var user User
	err := config.DB.Where("id = ?", id).First(&user)
	if err.Error != nil {
		return (User{}), err.Error
	}
	return user, nil
}

func (user *User) DeleteByIdService() {
	config.DB.Delete(&User{}, user.ID)
}

func (user *User) UpdateByIdService(updatedUser User) {
	config.DB.Where("id = ?", user.ID).Updates(updatedUser)
}

func FindAndDeleteByIdService(id int) error {
	user, err := FindByIDService(id)
	if err != nil {
		return err
	}
	user.DeleteByIdService()
	return nil
}

func FindAndUpdateByIdService(id int, updatedUserBody User) (User, error) {
	user, err := FindByIDService(id)
	if err != nil {
		return (User{}), err
	}
	user.UpdateByIdService(updatedUserBody)
	updatedUser, _ := FindByIDService(id)
	return updatedUser, nil
}
