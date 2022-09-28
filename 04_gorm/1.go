package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type User struct {
	Id       uint      `gorm:"column:id;PRIMARY_KEY;NOT NULL;AUTO_INCREMENT;"`
	Name     string    `gorm:"column:name"`
	Age      int64     `gorm:"column:age"`
	Birthday time.Time `gorm:"column:birthday"`
	Email    string    `gorm:"default:'123'"`
	Address  string    `gorm:"column:address"`
	Info     string    `gorm:"-"` //忽略这个字段
}

func main() {
	//gorm.Open("mysql", "user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local")
	db, _ := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/gorm_test_local?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Error)})

	//**插入*****************************************************************************

	user := User{Name: "胡适", Age: 43, Birthday: time.Now(), Email: "32131231", Address: "南山区"}
	//插入一行数据
	db.Create(&user)

	//**查询*****************************************************************************

	//查询第一行
	first := User{} //定义接收查询结果的结构体变量
	db.First(&first)
	fmt.Println("查询第一行:", first)

	//查询一条记录，默认是第一条
	take := User{}
	db.Take(&take)
	fmt.Println("查询一条记录，默认是第一条:", take)

	//查询最后一条记录
	last := User{}
	db.Last(&last)
	fmt.Println("查询最后一条记录:", last)

	//查询所有记录
	fmt.Println()
	var find []User
	db.Find(&find)
	//fmt.Println("查询所有记录：", find)

	//查询一列值
	fmt.Println()
	var pluck []string
	db.Model(&User{}).Pluck("name", &pluck)
	//fmt.Println("查询一列值:", pluck)

	//在某一范围内查询一个
	var where []User
	db.Where("age in (?)", []int{43, 41}).Take(&where)
	fmt.Println("在某一范围内查询一个：", where)

	//查询内范围
	var find1 []User
	db.Where("age >= ? and age <= ?", 40, 42).Find(&find1)
	fmt.Println("查询内范围：", find1)

	//limit
	var limit []*User
	db = db.Limit(3).Where("address = ?", "南山区").Offset(3).Find(&limit)
	fmt.Println("limit: ", limit)
	fmt.Println(len(limit))
	fmt.Println(*limit[0])

	//**更新*****************************************************************************
	/*
		update := Users{}
		db.Where("id = ?", 5).Take(&update)
		update.Age = 100
		db.Save(&update)

		//只更新id为5的那一行
		//db.Model(&update).Update("birthday", time.Now())

		updateUsers1 := User{
			Age:      200,
			Name:     "李大钊",
			Birthday: time.Now(),
		}
		db.Model(&update).Updates(&updateUsers1)

		//更新所有行
		//db.Model(&User{}).Update("birthday", time.Now())

		//删除数据
		db.Where("id = ?", 26).Delete(&Users{})*/

}
