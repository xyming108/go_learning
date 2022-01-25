package model

import (
	"GO_Learning/13_project/01_go_excel_batch/utils/modelUtils"
	"fmt"
)

type OtherSnMac struct {
	ProductCode string `gorm:"column:product_code"`
	Sn          string `gorm:"column:sn"`
	Mac         string `gorm:"column:mac"`
}

func (x *OtherSnMac) TableName() string {
	return "other_sn_mac"
}

func InsertBatch(temp [][]string, length int) error {
	db := modelUtils.GetDB()

	sqlStr := "INSERT INTO other_sn_mac(product_code, sn, mac) VALUES"
	for i := 0; i < length; i++ { // 批量插入
		if i == 0 {
			sqlStr += fmt.Sprintf("('%s', '%s', '%s')", temp[i][0], temp[i][1], temp[i][2])
		} else {
			sqlStr += fmt.Sprintf(",('%s', '%s', '%s')", temp[i][0], temp[i][1], temp[i][2])
		}
	}

	if err := db.Exec(sqlStr).Error; err != nil {
		return err
	}

	return nil
}
