package main

/*import "fmt"

type Phone interface {
	Show()
}

type AbstractFactory interface {
	ProducePhone() Phone
}

// 手机模块
type Xiaomi struct {

}

func (x *Xiaomi) Show()  {
	fmt.Println("xiaomi")
}

type Huawei struct {

}

func (h *Huawei) Show()  {
	fmt.Println("huawei")
}

// 手机工厂模块
type XiaomiFactory struct {

}

func (x *XiaomiFactory) ProducePhone() Phone {
	return new(Xiaomi)
}

type HuaweiFactory struct {

}

func (h *HuaweiFactory) ProducePhone() Phone {
	return new(Huawei)
}

func FactoryMethod01()  {
	f1 := new(XiaomiFactory)
	f1.ProducePhone().Show() // 生产小米手机

	f2 := new(HuaweiFactory)
	f2.ProducePhone().Show() // 生产华为手机
}
*/