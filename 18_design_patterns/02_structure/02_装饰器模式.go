package _2_structure

import "fmt"

// 抽象产品接口
type product interface {
	getPrice() float64
}

// 具体产品
type concreteProduct struct {
	price float64
}

func (c *concreteProduct) getPrice() float64 {
	c.price = 10.0
	return c.price
}

// 添加具体装饰类，涨价
type productAddDecorator struct {
	product product
}

func (p *productAddDecorator) getPrice() float64 {
	return p.product.getPrice() + 3.0
}

// 添加具体装饰类，降价
type productDelDecorator struct {
	product product
}

func (p *productDelDecorator) getPrice() float64 {
	return p.product.getPrice() - 2.0
}

func ProxyFunc02() {
	p := new(concreteProduct)
	fmt.Println("原价：", p.getPrice())

	add := &productAddDecorator{product: p}
	fmt.Println("涨价后：", add.getPrice())

	del := &productDelDecorator{product: add}
	fmt.Println("降价后：", del.getPrice())
}
