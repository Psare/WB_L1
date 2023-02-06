package main

import "fmt"

// Есть клиент, ожидающий определенных качеств (порт Lightning),
// но также есть объект adapter (ноутбук на Windows), который предоставляет тот же функционал,
// но через другой интерфейс (USB порт).
// Создаем структуру adapter, которая будет:
// Реализует тот же интерфейс, который ожидает клиент (порт Lightning);
// Переводить запрос от клиента к адаптируемому объекту в форме, которую он ожидает.
// Адаптер принимает коннектор Lightning, после чего переводит его сигналы в формат USB в ноутбуке на Windows.

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