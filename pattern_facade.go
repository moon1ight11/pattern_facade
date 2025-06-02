package main

import (
	"fmt"
	"time"
	"math/rand"
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


// Подсистема 6: Режим кпп
type transmission_mode struct {}

func (tm *transmission_mode) set_trans_mode(transmission_mode string) {
	fmt.Printf("Трансмиссия установлена в режим %s\n", transmission_mode)
}


// Подсистема 7: Охрана
type sequrity struct {}

func (sq *sequrity) Off() {
	fmt.Println("Сигнализация деактивирована")
}
func (sq *sequrity) On() {
	fmt.Println("Сигнализация активирована")
}
func (sq *sequrity) Heating() {
	heat_time := rand.Intn(15)

	fmt.Println("Попытка запуска ДВС...")
	time.Sleep(4 * time.Second)
	fmt.Printf("Двигатель запущен, до полного прогрева примерно %d минут \n", heat_time)
}


// Фасад: умный авто
type SmartAuto struct {
	lights *lights
	hand_brake *hand_brake
	selector *selector
	speed_limiter *speed_limiter
	conditioner *conditioner
	transmission_mode *transmission_mode
	sequrity *sequrity
}


// Конструктор фасада
func new_SmartAuto() *SmartAuto {
	return &SmartAuto{
		lights: &lights{},
		hand_brake: &hand_brake{},
		selector: &selector{},
		speed_limiter: &speed_limiter{},
		conditioner: &conditioner{},
		transmission_mode: &transmission_mode{},
		sequrity: &sequrity{},
	}
}


// Метод для включения режима "Стоянка"
func (s *SmartAuto) ParkingMode() {
	fmt.Println("Активация режима 'Стоянка'")
	time.Sleep(2 * time.Second)
	s.lights.Off()
	time.Sleep(1 * time.Second)
	s.hand_brake.On()
	time.Sleep(1 * time.Second)
	s.selector.Parking()
	time.Sleep(1 * time.Second)
	s.speed_limiter.Off()
	time.Sleep(1 * time.Second)
	s.conditioner.Off()
	time.Sleep(1 * time.Second)
	s.sequrity.On()
	time.Sleep(2 * time.Second)
	fmt.Println("Режим 'Стоянка' успешно активирован")
}


// Метод для включения режима "Поездка"
func (s *SmartAuto) DriveMode() {
	fmt.Println("Активация режима 'Поездка'")
	time.Sleep(2 * time.Second)

	var limit int
	fmt.Println("Установите ограничение скорости")
	fmt.Scan(&limit)
	if limit > 180 {
		fmt.Println("Ошибка. Указанное значение установить невозможно")
		return
	}
	
	var temp int
	fmt.Println("Установите температуру воздуха в салоне")
	fmt.Scan(&temp)
	if temp < 10 && temp > 35 {
		fmt.Println("Ошибка. Температура вне диапазона работы кондиционера")
		return
	}

	var transmission_mode string
	fmt.Println("Выберите режим работы КПП (Sport/Dirt/Eco)")
	fmt.Scan(&transmission_mode)
	if transmission_mode != "Sport" && transmission_mode != "Dirt" && transmission_mode != "Eco" {
		fmt.Println("Ошибка. Некорректно указан режим")
		return
	}
	time.Sleep(2 * time.Second)

	s.sequrity.Off()
	time.Sleep(1 * time.Second)
	s.lights.On()
	time.Sleep(1 * time.Second)
	s.hand_brake.Off()
	time.Sleep(1 * time.Second)
	s.selector.Drive()
	time.Sleep(1 * time.Second)
	s.transmission_mode.set_trans_mode(transmission_mode)
	time.Sleep(1 * time.Second)
	s.speed_limiter.set_speed_limit(limit)
	time.Sleep(1 * time.Second)
	s.conditioner.set_temp_conditioner(temp)
	time.Sleep(2 * time.Second)
	

	fmt.Println("Режим 'Поездка' успешно активирован")
}


// метод для режима Прогрев
func (s *SmartAuto) HeatingMode() {
	fmt.Println("Активация режима 'Прогрев'")
	time.Sleep(2 * time.Second)

	s.lights.On()
	time.Sleep(1 * time.Second)
	s.hand_brake.On()
	time.Sleep(1 * time.Second)
	s.selector.Parking()
	time.Sleep(1 * time.Second)
	s.sequrity.Heating()
	time.Sleep(2 * time.Second)

	fmt.Println("Режим 'Прогрев' успешно активирован")
}



func main() {
	smartCar := new_SmartAuto()

	var Mode string
	fmt.Println("Какой режим нужно активировать?(Drive/Parking/Heating)")
	fmt.Scan(&Mode)

	if Mode == "Drive" {
		smartCar.DriveMode()
	} else if Mode == "Parking" {
		smartCar.ParkingMode()
	} else if Mode == "Heating" {
		smartCar.HeatingMode()
	} else {
		fmt.Println("Некорректно указан режим")
		return
	}
}