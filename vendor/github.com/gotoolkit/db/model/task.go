package model

import (
	"github.com/jinzhu/gorm"
	"github.com/gotoolkit/db/orm"
	"time"
)

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Priority  string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `json:"done"`
}

func (u *Task) GetAll() (interface{}, error) {
	users := []Task{}
	err := orm.GetDB().Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *Task) GetByID(id int) (interface{}, error) {
	task := new(Task)
	err := orm.GetDB().Find(task, id).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *Task) Create(in interface{}) error {
	err := orm.GetDB().Save(in).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *Task) BindModel() interface{} {
	return new(Task)
}
