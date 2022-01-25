package service

import (
	"GO_Learning/13_project/01_go_excel_batch/model"
	"log"
)

func ChangeState(id int) error {
	err := model.UpdateState(id)
	if err != nil {
		log.Println("update state err: ", err)
		return err
	}
	return nil
}
