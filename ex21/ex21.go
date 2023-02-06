package main

import "fmt"

// Этот код представляет реализацию паттерна "адаптер" в Go.
// Имеется клиент, у которого есть метод insertLightningConnectorIntoComputer, который ожидает в качестве аргумента объект, 
// реализующий интерфейс computer. Этот интерфейс определяет метод insertIntoLightningPort.
// Существует структура mac, реализующая интерфейс computer и имеющая метод insertIntoLightningPort. 
// Она может быть использована клиентом без проблем.
// Также существует структура windows, которая не реализует интерфейс computer, но имеет метод insertIntoUSBPort. 
// Для использования этой структуры с клиентом необходимо использовать адаптер.
// Адаптер реализует интерфейс computer и использует структуру windows внутри себя. Когда клиент вызывает метод insertIntoLightningPort, 
// адаптер преобразует сигнал и вызывает метод insertIntoUSBPort у структуры windows.
// В main создаются экземпляры клиента, mac и windows.

// Клиентский код.
type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

// Интерфейс клиента.
type computer interface {
	insertIntoLightningPort()
}

// Сервис.
type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Неизвестный сервис.
type windows struct{}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Адаптер.
type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}