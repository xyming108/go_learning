package _2_structure

import "fmt"

// 买房抽象接口
type BuyHouse interface {
	Look()
	Talk()
	Buy()
}

// 具体买房者
type Buyer struct {
	Name string
}

func (b *Buyer) Look() {
	fmt.Println(b.Name, " 看房子")
}

func (b *Buyer) Talk() {
	fmt.Println(b.Name, " 和房东交谈")
}

func (b *Buyer) Buy() {
	fmt.Println(b.Name, " 决定买房子")
}

// 代理买房者（中介）
type ProxyBuyer struct {
	Buyer *Buyer
}

func (p *ProxyBuyer) Look() {
	fmt.Printf("看房前中介联系了 ")
	p.Buyer.Look()
}

func (p *ProxyBuyer) Talk() {
	fmt.Printf("中介建议 ")
	p.Buyer.Talk()
}

func (p *ProxyBuyer) Buy() {
	fmt.Printf("中介得知 ")
	p.Buyer.Buy()
}

func NewProxyBuyer(name string) *ProxyBuyer {
	return &ProxyBuyer{
		Buyer: &Buyer{
			Name: name,
		},
	}
}

func ProxyFunc01() {
	proxy := NewProxyBuyer("小明")
	proxy.Look()
	proxy.Talk()
	proxy.Buy()
}
