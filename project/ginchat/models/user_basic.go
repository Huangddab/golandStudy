package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"column:name"`
	PassWord      string    `gorm:"column:pass_word"`
	Email         string    `gorm:"type:varchar(100);unique"`
	Identity      string    `gorm:"column:identity;type:varchar(32)"`
	ClientIp      string    `gorm:"column:client_ip" json:"client_ip"`
	ClientPort    string    `gorm:"column:client_port" json:"client_port"`
	LoginTime     time.Time `gorm:"column:login_time" json:"login_time"`
	HeartBeatTime time.Time `gorm:"column:heart_beat_time" json:"heart_beat_time"`
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool      `gorm:"column:is_logout" json:"is_logout"`
	DeviceInfo    string    `gorm:"column:device_info" json:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
