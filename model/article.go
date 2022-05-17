package model

import (
	"blogo/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid""`
	gorm.Model
	Title   string `gorm:"type:varchar(100); not null"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:""desc`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ErrArtNotExist
	}
	return art, errmsg.SUCCESS
}

func GetArtForCate(id, pageSize, pageNum int) ([]Article, int) {
	var arts []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", id).Find(&arts).Error
	if err != nil {
		return nil, errmsg.ErrCateNameNotExist
	}
	return arts, errmsg.SUCCESS
}

func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ListArts(pageSize, pageNum int) ([]Article, int) {
	var arts []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return arts, errmsg.SUCCESS
}

func UpdateArt(id int, art *Article) int {
	var a Article
	var maps = make(map[string]interface{})
	maps["title"] = art.Title
	maps["cid"] = art.Cid
	maps["desc"] = art.Desc
	maps["content"] = art.Content
	maps["img"] = art.Img
	err := db.Model(&a).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Unscoped().Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.ERROR
}
