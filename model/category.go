package model

import (
	"blogo/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCate(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ErrCateNameUsed
	}
	return errmsg.SUCCESS
}

func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ListCate(pageSize, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

func UpdateCate(id int, cate *Category) int {
	var c Category
	var maps = make(map[string]interface{})
	maps["name"] = cate.Name
	err := db.Model(&c).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Unscoped().Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.ERROR
}
