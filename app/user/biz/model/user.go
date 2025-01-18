package model

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/utils"
	"gorm.io/datatypes"
	"time"
)

type User struct {
	PublicModel
	Email       string         `json:"email" gorm:"column:email"`
	Username    string         `json:"username" gorm:"column:username"`
	Openid      string         `json:"openid" gorm:"column:openid"`
	Salt        string         `json:"salt" gorm:"column:salt"`
	Password    string         `json:"password" gorm:"column:password"`
	Photo       datatypes.JSON `json:"photo" gorm:"column:photo"`
	Gender      int8           `json:"gender" gorm:"column:gender"`
	BirthDate   time.Time      `json:"birth_date" gorm:"column:birth_date"`
	IsVisible   int8           `json:"is_visible" gorm:"column:is_visible"`
	Signature   string         `json:"signature" gorm:"column:signature"`
	SocialMedia string         `json:"social_media" gorm:"social_media"`
}

type UserList []User

func (User) TableName() string {
	return "lv_users"
}

// IsExistByField , judge if user exist by field
func (user *User) IsExistByField(field string, value any) bool {
	if dal.DB == nil {
		dal.Init()
	}
	if err := dal.DB.Where(field, value).Find(&user).Error; err != nil {
		return false
	}
	if user.ID <= 0 {
		return false
	}
	return true
}

// Create , insert user into database
func (user *User) Create() bool {
	if err := dal.DB.Create(&user).Error; err != nil {
		return false
	}
	return true
}

// ValidatePassword , check if input password is correct
func (user *User) ValidatePassword(email, password string) bool {
	return utils.EncodePassword(email, password) == user.Password
}
