package model

import "GO_Learning/13_project/01_go_excel_batch/utils/modelUtils"

type UploadSnMac struct {
	Id         int    `gorm:"primary_key;column:id"`
	Name       string `gorm:"column:name"`
	Url        string `gorm:"column:url"`
	Size       int    `gorm:"column:size"`
	State      int    `gorm:"column:state"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (x *UploadSnMac) TableName() string {
	return "01_go_excel_batch"
}

func UpdateState(id int) error {
	db := modelUtils.GetDB()

	var uploadSnMac UploadSnMac

	err := db.Model(&uploadSnMac).Where("id = ?", id).Update("state", 2).Error
	if err != nil {
		return err
	}
	return nil
}
