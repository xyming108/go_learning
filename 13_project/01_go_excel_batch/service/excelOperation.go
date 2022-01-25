package service

import (
	"GO_Learning/13_project/01_go_excel_batch/model"
	"GO_Learning/13_project/01_go_excel_batch/types"
	"GO_Learning/13_project/01_go_excel_batch/utils/csvUtils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func ExcelOperation(data types.TransferData) {
	Id := data.Id
	url := data.Url
	name := data.Name

	res, err := http.Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()

	fileName := fmt.Sprintf("%v.%s", time.Now().Format("2006-01-02_15:04:05"), name)
	path := fmt.Sprintf("temp/%s", fileName)

	localFile, err := os.Create(path)
	if err != nil {
		return
	}

	defer localFile.Close()

	_, err = io.Copy(localFile, res.Body)
	if err != nil {
		return
	}

	//并发，批处理
	err = ExcelBatchProcess(path)
	if err != nil {
		return
	}

	//更改状态
	err = ChangeState(Id)
	if err != nil {
		return
	}

	//移除临时文件
	err = os.Remove(path)
}

func ExcelBatchProcess(path string) error {
	_, data, err := csvUtils.CsvReadAll(path)
	if err != nil {
		return err
	}

	//每100条数据为一组批量插入数据
	getInt := len(data) / 100 //取整
	getMod := len(data) % 100 //取模

	tempData := make(map[int][][]string, len(data))
	if len(data) < 100 { //情况一：excel数据总行数小于100，不用go并发，直接批量插入
		if getMod == 0 {
			log.Print("数据为空")
			return err
		}
		for i := 0; i < len(data); i++ {
			tempData[0] = append(tempData[0], data[i])
		}

		//开始批量插入
		err := model.InsertBatch(tempData[0], len(data))
		if err != nil {
			log.Println("db insert err: ", err)
			return err
		}
	} else if len(data) >= 100 { //情况二：excel数据总行数大于100
		//情况二里面的第一种情况：excel数据总行数大于100，数据条数正好可以被100整除
		var i, j int
		for i = 0; i < getInt; {
			tempData[i] = append(tempData[i], data[j])
			j++
			if j%100 == 0 {
				i++
			}
		}

		//在所有批处理数据组中，再次分组，每50组开启一个goroutine
		getIntSecond := getInt / 50 //取整
		getModSecond := getInt % 50 //取模

		if getIntSecond == 0 && getModSecond != 0 { //情况一：goroutine组的组数小于50
			s := make(chan [][]string, getModSecond)
			go func() {
				for j := 0; j < getModSecond; j++ {
					s <- tempData[j]
				}
				close(s)
			}()

			GoConcurrency(s)
		} else if getIntSecond != 0 { //情况二：goroutine组的组数大于50，且能被整除
			for i := 0; i < getIntSecond; i++ {
				s := make(chan [][]string, 50)
				go func() {
					for j := i; j < i+50; j++ {
						s <- tempData[j]
					}
					close(s)
				}()

				GoConcurrency(s)
			}

			if getModSecond != 0 { //处理余下不能被整除的goroutine组
				s := make(chan [][]string, getModSecond)
				go func() {
					for j := getIntSecond * 50; j < getIntSecond*50+getModSecond; j++ {
						s <- tempData[j]
					}
					close(s)
				}()

				GoConcurrency(s)
			}
		}

		if getMod != 0 { //情况二里面的第二种情况：excel数据总行数大于100，数据总条数不能被100整除
			for k := getInt * 100; k < len(data); k++ {
				tempData[getInt] = append(tempData[getInt], data[k])
			}

			//开始批量插入末尾一组
			err := model.InsertBatch(tempData[getInt], getMod)
			if err != nil {
				log.Printf("db insert err: %s, err data: %v", err, tempData[getInt])
				return err
			}
		}

		log.Println("concurrency insert successfully")
	}

	return nil
}

func GoConcurrency(s chan [][]string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			z, ok := <-s
			if !ok {
				break
			}

			err := model.InsertBatch(z, 100)
			if err != nil {
				log.Println("db insert err: ", err)
				return
			}
		}
	}()

	wg.Wait()
}
