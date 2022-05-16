package model

import (
	"blogo/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ErrUsernameUsed
	}
	return errmsg.SUCCESS
}

func CreateUser(user *User) int {
	// user.Password = ScryptPw(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ListUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

func UpdateUser(id int, user *User) int {
	var u User
	var maps = make(map[string]interface{})
	maps["username"] = user.UserName
	maps["role"] = user.Role
	err := db.Model(&u).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Unscoped().Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.ERROR
}

func ScryptPw(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 88, 98, 34, 123, 255, 192}
	key, err := scrypt.Key([]byte(password), salt, 16364, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	pwd := base64.StdEncoding.EncodeToString(key)
	return pwd

}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}
