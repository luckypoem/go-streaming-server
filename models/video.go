package models

import "go-streaming-server/orm"

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
