package _2_structure

import "fmt"

// 挂号
type registration struct{}

func (r *registration) Registration() {
	fmt.Println("挂号")
}

// 门诊
type outpatientService struct{}

func (o *outpatientService) OutpatientService() {
	fmt.Println("门诊")
}

// 付费
type pricing struct{}

func (p *pricing) Pricing() {
	fmt.Println("付费")
}

// 取药
type medicationCollection struct{}

func (m *medicationCollection) MedicationCollection() {
	fmt.Println("取药")
}

// 接待人员
type receptionist struct {
	registration         registration
	outpatientService    outpatientService
	pricing              pricing
	medicationCollection medicationCollection
}

// 想去看病的流程
func (r *receptionist) FirstVisit() {
	r.registration.Registration()
	r.outpatientService.OutpatientService()
	r.pricing.Pricing()
	r.medicationCollection.MedicationCollection()
}

// 看完了只想想拿药的流程
func (r *receptionist) MedicationCollection() {
	r.medicationCollection.MedicationCollection()
}

func ProxyFunc04() {
	// 情况一
	receptionist := receptionist{}
	fmt.Println("1. 想去看病的流程: ")
	receptionist.FirstVisit() // 挂号 门诊 付费 取药

	// 情况二
	fmt.Println("2. 看完了只想想拿药的流程: ")
	receptionist.MedicationCollection() // 只取药
}
