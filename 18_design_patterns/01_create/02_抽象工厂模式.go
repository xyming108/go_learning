package main

import "fmt"

// 抽象层
type AbstractGPU interface {
	Display()
}

type AbstractCPU interface {
	Calculate()
}

type AbstractMemory interface {
	Storage()
}

// 抽象工厂
type AbstractFactory interface {
	CreateGPU() AbstractGPU
	CreateCPU() AbstractCPU
	CreateMemory() AbstractMemory
}

// Intel 厂商
type IntelGPU struct{}
type IntelCPU struct{}
type IntelMemory struct{}
type IntelFactory struct{}

func (i *IntelGPU) Display() {
	fmt.Printf("Intel GPU display，")
}

func (i *IntelCPU) Calculate() {
	fmt.Printf("Intel CPU calculate，")
}

func (i *IntelMemory) Storage() {
	fmt.Printf("Intel memory storage，")
}

func (i *IntelFactory) CreateIntelGPU() AbstractGPU {
	return new(IntelGPU)
}

func (i *IntelFactory) CreateIntelCPU() AbstractCPU {
	return new(IntelCPU)
}

func (i *IntelFactory) CreateIntelMemory() AbstractMemory {
	return new(IntelMemory)
}

// Nvidia 厂商
type NvidiaGPU struct{}
type NvidiaCPU struct{}
type NvidiaMemory struct{}
type NvidiaFactory struct{}

func (n *NvidiaGPU) Display() {
	fmt.Printf("Nvidia GPU display，")
}

func (n *NvidiaCPU) Calculate() {
	fmt.Printf("Nvidia CPU calculate，")
}

func (n *NvidiaMemory) Storage() {
	fmt.Printf("Nvidia memory storage，")
}

func (n *NvidiaFactory) CreateNvidiaGPU() AbstractGPU {
	return new(NvidiaGPU)
}

func (n *NvidiaFactory) CreateNvidiaCPU() AbstractCPU {
	return new(NvidiaCPU)
}

func (n *NvidiaFactory) CreateNvidiaMemory() AbstractMemory {
	return new(NvidiaMemory)
}

// Kingston 厂商
type KingstonGPU struct{}
type KingstonCPU struct{}
type KingstonMemory struct{}
type KingstonFactory struct{}

func (k *KingstonGPU) Display() {
	fmt.Printf("Kingston GPU display，")
}

func (k *KingstonCPU) Calculate() {
	fmt.Printf("Kingston CPU calculate，")
}

func (k *KingstonMemory) Storage() {
	fmt.Printf("Kingston memory storage，")
}

func (k *KingstonFactory) CreateKingstonGPU() AbstractGPU {
	return new(KingstonGPU)
}

func (k *KingstonFactory) CreateKingstonCPU() AbstractCPU {
	return new(KingstonCPU)
}

func (k *KingstonFactory) CreateKingstonMemory() AbstractMemory {
	return new(KingstonMemory)
}

func FactoryMethod02() {
	// Intel的CPU，Intel的显卡，Intel的内存
	intelFac := new(IntelFactory)
	intelFac.CreateIntelCPU().Calculate()
	intelFac.CreateIntelGPU().Display()
	intelFac.CreateIntelMemory().Storage()
	fmt.Println()

	// Intel的CPU， nvidia的显卡，Kingston的内存
	nvidiaFac := new(NvidiaFactory)
	kingstonFac := new(KingstonFactory)
	intelFac.CreateIntelCPU().Calculate()
	nvidiaFac.CreateNvidiaGPU().Display()
	kingstonFac.CreateKingstonMemory().Storage()
}
