package main

import (
	"fmt"
)

// Подсистема 1: свет
type lights struct{}

func (l *lights) Off() {
	fmt.Println("Свет выключен")
}
func (l *lights) On() {
	fmt.Println("Свет включен")
}

// Подсистема 2: ручной тормоз
type hand_brake struct {}

func (h *hand_brake) Off() {
	fmt.Println("Ручной тормоз снят")
}
func (h *hand_brake) On() {
	fmt.Println("Ручной тормоз поставлен")
}

// Подсистема 3: положение кпп
type selector struct {} 

func (s *selector) Drive() {
	fmt.Println("Селектор КПП в положении DRIVE")
}
func (s *selector) Parking() {
	fmt.Println("Селектор КПП в положении PARKING")
}

// Фасад: умный авто
type SmartAuto struct {
	lights *lights
	hand_brake *hand_brake
	selector *selector
}

// Конструктор фасада
func new_SmartAuto() *SmartAuto {
	return &SmartAuto{
		lights: &lights{},
		hand_brake: &hand_brake{},
		selector: &selector{},
	}
}

// Метод для включения режима "Стоянка"
func (s *SmartAuto) ParkingMode() {
	fmt.Println("Активация режима Стоянка")
	s.lights.Off()
	s.hand_brake.On()
	s.selector.Parking()
	fmt.Println("Режим Стоянка активирован")
}

// Метод для включения режима "Поездка"
func (s *SmartAuto) DriveMode() {
	fmt.Println("Активация режима Поездка")
	s.lights.On()
	s.hand_brake.Off()
	s.selector.Drive()
	fmt.Println("Режим Поездка активирован")
}

func main() {
	smartCar := new_SmartAuto()

	var Mode string
	fmt.Println("Какой режим нужно активировать?")
	fmt.Scan(&Mode)

	if Mode == "Drive" {
		smartCar.DriveMode()
	} else if Mode == "Parking" {
		smartCar.ParkingMode()
	}
}