package _2_structure

import "fmt"

// 客户端
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("客户使用lightning接口")
	com.InsertIntoLightningPort()
}

// Computer
type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("lightning接口连接到mac")
}

type Windows struct {
}

func (w *Windows) InsertIntoUsbPort() {
	fmt.Println("usb接口连接到windows")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("适配器把lightning接口适配成use接口")
	w.windowMachine.InsertIntoUsbPort()
}

func ProxyFunc03() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
