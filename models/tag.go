package models

import (
	"github.com/jinzhu/gorm"
	"time"
)


// 在同个models包下，因此db *gorm.DB是可以直接使用的

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}


func (t *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}


func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag, err error) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID >0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return  true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}


func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on !=?",0).Delete(&Tag{})
	return true
}