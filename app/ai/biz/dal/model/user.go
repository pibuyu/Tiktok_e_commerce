package model

import "C"
import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id             int32
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (result *User, err error) {
	err = db.Model(&User{}).Where("email = ?", email).First(&result).Error
	return
}

func GetById(id int32) (result *User, err error) {
	err = mysql.DB.Model(&User{}).Where("id = ?", id).First(&result).Error
	return
}

func DeleteUserById(db *gorm.DB, userId int32) error {
	return db.Model(&User{}).Where("id = ?", userId).Delete(&User{}).Error
}

func UpdatePassword(db *gorm.DB, email, newPassword string) error {
	return db.Model(&User{}).Where("email = ?", email).UpdateColumn("password_hashed", newPassword).Error
}
