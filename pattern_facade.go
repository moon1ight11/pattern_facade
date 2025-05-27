package main

import (
	"fmt"
)

// Подсистема 1: Ходовые огни
type lights struct{}

func (l *lights) Off() {
	fmt.Println("Ходовые огни выключены")
}
func (l *lights) On() {
	fmt.Println("Ходовые огни включены")
}

// Подсистема 2: Стояночный тормоз
type hand_brake struct {}

func (h *hand_brake) Off() {
	fmt.Println("Стояночный тормоз снят")
}
func (h *hand_brake) On() {
	fmt.Println("Стояночный тормоз установлен")
}

// Подсистема 3: Положение селектора кпп
type selector struct {} 

func (s *selector) Drive() {
	fmt.Println("Селектор КПП в положении DRIVE")
}
func (s *selector) Parking() {
	fmt.Println("Селектор КПП в положении PARKING")
}

// Подсистема 4: Ограничение скорости
type speed_limiter struct {}

func (sl *speed_limiter) Off() {
	fmt.Println("Ограничитель скорости выключен")
}
func (sl *speed_limiter) set_speed_limit(limit int) {
	fmt.Printf("Ограничитель скорости: установлено ограничение в %d км/ч\n", limit)
}

// Подсистема 5: Кондиционер
type conditioner struct {}

func (c *conditioner) Off() {
	fmt.Println("Кондиционер выключен")
}
func (c *conditioner) set_temp_conditioner(temp int) {
	fmt.Printf("Кондиционер: установлена температура %d°C\n", temp)
}

// Фасад: умный авто
type SmartAuto struct {
	lights *lights
	hand_brake *hand_brake
	selector *selector
	speed_limiter *speed_limiter
	conditioner *conditioner
}

// Конструктор фасада
func new_SmartAuto() *SmartAuto {
	return &SmartAuto{
		lights: &lights{},
		hand_brake: &hand_brake{},
		selector: &selector{},
		speed_limiter: &speed_limiter{},
		conditioner: &conditioner{},
	}
}

// Метод для включения режима "Стоянка"
func (s *SmartAuto) ParkingMode() {
	fmt.Println("Активация режима 'Стоянка'")
	s.lights.Off()
	s.hand_brake.On()
	s.selector.Parking()
	s.speed_limiter.Off()
	s.conditioner.Off()
	fmt.Println("Режим 'Стоянка' успешно активирован")
}

// Метод для включения режима "Поездка"
func (s *SmartAuto) DriveMode() {
	fmt.Println("Активация режима 'Поездка'")

	var limit int
	fmt.Println("Установите ограничение скорости")
	fmt.Scan(&limit)
	
	var temp int
	fmt.Println("Установите температуру воздуха в салоне")
	fmt.Scan(&temp)

	s.lights.On()
	s.hand_brake.Off()
	s.selector.Drive()
	s.speed_limiter.set_speed_limit(limit)
	s.conditioner.set_temp_conditioner(temp)

	fmt.Println("Режим 'Поездка' успешно активирован")
}

func main() {
	smartCar := new_SmartAuto()

	var Mode string
	fmt.Println("Какой режим нужно активировать?(Drive/Parking)")
	fmt.Scan(&Mode)

	if Mode == "Drive" {
		smartCar.DriveMode()
	} else if Mode == "Parking" {
		smartCar.ParkingMode()
	}
}