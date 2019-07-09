package models

import "go-streaming-server/orm"

type Video struct {
	Id     int
	FileId string
	Token  string
}

func init() {
	orm.Gorm.AutoMigrate(new(Video))
}
