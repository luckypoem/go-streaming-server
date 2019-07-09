package models

import (
	"go-streaming-server/orm"
	"strconv"
)

type Video struct {
	Id     int
	FileId int
	Token  string
}

func init() {
	orm.Gorm.AutoMigrate(new(Video))
}

func CreateVideo(v *Video) error {
	err := orm.Gorm.Create(v).Error

	if err != nil {
		return err
	}

	return nil
}

func CheckVideo(fileid int) bool {
	var videos []*Video

	err := orm.Gorm.Where("file_id = ?", strconv.Itoa(fileid)).Find(&videos).Error

	if err != nil {
		return false
	}

	if len(videos) != 1 {
		return false
	}

	return true
}

func CheckToken(fileid int, token string) bool {
	var videos []*Video

	err := orm.Gorm.Where("file_id = ?", strconv.Itoa(fileid)).Where("token = ?", token).Find(&videos).Error

	if err != nil {
		return false
	}

	if len(videos) != 1 {
		return false
	}

	return true
}
