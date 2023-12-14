package main

import (
	"fmt"
)

// Car является целевым интерфейсом, который ожидает клиентский код.
type Car interface {
	GetTypesTransmission() string
}

// Mazda представляет адаптер, который преобразует интерфейс Transmission в интерфейс Car.
type Mazda struct {
	Transmission *Transmission
}

// Request вызывает адаптированный метод SpecificRequest на объекте Transmission.
func (a *Mazda) GetTypesTransmission() string {
	return a.Transmission.GetTypes()
}

// Transmission представляет собой существующий класс, который несовместим с интерфейсом Car.
type Transmission struct {
	Types []string
}

// SpecificRequest представляет собой специфический метод, который нужно адаптировать.
func (a *Transmission) GetTypes() string {
	result := "The following types of gearboxes are available:\n"
	for _, elem := range a.Types {
		result += elem + "\n"
	}
	return result
}

func main() {
	transmission := &Transmission{Types: []string{"Mechanical", "Automatic", "Robotic", "Variable"}}
	adapter := &Mazda{Transmission: transmission}

	result := adapter.GetTypesTransmission()
	fmt.Println(result)
}
