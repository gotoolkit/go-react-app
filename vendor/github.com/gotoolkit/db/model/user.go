package model

import (
	"github.com/jinzhu/gorm"
	"github.com/gotoolkit/db/orm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) GetAll() (interface{}, error) {
	users := []User{}
	err := orm.GetDB().Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *User) GetByID(id int) (interface{}, error) {
	user := new(User)
	err := orm.GetDB().Find(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Create(in interface{}) error {
	err := orm.GetDB().Save(in).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BindModel() interface {} {
	return new(User)
}