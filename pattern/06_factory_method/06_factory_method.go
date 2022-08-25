package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
[1] https://golangbyexample.com/golang-factory-design-pattern/
[2] https://golang-blog.blogspot.com/2019/10/factory-pattern.html
[3] https://github.com/AlexanderGrom/go-patterns/tree/master/Creational/FactoryMethod

Создаем игровкю приставку
*/

// [0] интерфейс приставки
type iGameConsole interface {
	setName(name string)
	setCPU(cpu string)
	setGPU(cpu string)
	setHDD(hdd int)
	getConfig()
}

//[1] реализуем интерфейс структурой
type GameConsole struct {
	name string
	cpu  string
	gpu  string
	hdd  int
}

func (gc *GameConsole) setName(n string) {
	gc.name = n
}
func (gc *GameConsole) setCPU(cpu string) {
	gc.cpu = cpu
}
func (gc *GameConsole) setGPU(gpu string) {
	gc.gpu = gpu
}
func (gc *GameConsole) setHDD(mem int) {
	gc.hdd = mem
}
func (gc *GameConsole) getConfig() {
	fmt.Printf("[Config of %s]:\n CPU %s,\n GPU %s,\n MEM %s\n", gc.name, gc.cpu, gc.gpu, gc.hdd)
	fmt.Println("=======================================================================")
}

//==================================================

//[2] - реализуем конкретную консуоль - PS5
type PlayStation5 struct {
	GameConsole
}

// Возвращаем конфигурацию обьект
func newPS5() iGameConsole {
	return &PlayStation5{
		GameConsole: GameConsole{
			name: "Sony Playstation 5",
			cpu:  "Модифицированный AMD Ryzen 3-го поколения на базе 7-нм микроархитектуры Zen 2 (8-ядер/16-потоков) @ 3,5 ГГц",
			gpu:  "Модифицированный AMD Radeon Navi на базе 7-нм микроархитектуры RDNA 2 с поддержкой трассировки лучей @ 10,3 TFLOPS",
			hdd:  800,
		},
	}
}

// [3] реализуем xbox
type XBox struct {
	GameConsole
}

// Возвращаем конфигурацию обьект
func newXBox() iGameConsole {
	return &XBox{
		GameConsole: GameConsole{
			name: "Xbox Series X",
			cpu:  "Модифицированный AMD Ryzen 3-го поколения на базе 7-нм+ микроархитектуры Zen 2 (8-ядер/16-потоков) @ 3,8 ГГц, 3,6 ГГц с SMT",
			gpu:  "модифицированный AMD Radeon Navi на базе 7-нм+ микроархитектуры RDNA 2 с поддержкой трассировки лучей 52 CUs @ 1,825 ГГц, 12,2 TFLOPS",
			hdd:  1000,
		},
	}
}

// [4] фабрика приставок
func getConsole(consoleType string) (iGameConsole, error) {
	if consoleType == "XBox" {
		return newXBox(), nil
	}
	if consoleType == "PlayStation5" {
		return newPS5(), nil
	}
	return nil, fmt.Errorf("wrong console type")
}

// [5] - запускаем фабрику
func main() {
	ps5, _ := getConsole("PlayStation5")
	xbox, _ := getConsole("XBox")

	ps5.getConfig()
	xbox.getConfig()

}

/*
	Использование:
	1) когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код
	2) когда нужна возможность пользователям расширять части фреймворка или библиотеки
	3) когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых
	+:
	1) избавляет классы от привязки к конкретным классам продуктов
	2) выделяет код производства продуктов в одно место, упрощая поддержку кода
	3) упрощает добавление новых продуктов в программу
	4) реализует принцип открытости/закрытости
	-:
	1) может привести к созданию больших параллельных иерархий классов, так как для каждого класса
	продукта надо создать свой подкласс создателя
*/

// Плюсы
// 1. Избавляемся от конкретных типов структур и от инициализации с условиями, перенося это все в отдельный пакет, например
// 2. Код инициализации находится в одном месте
// 3. Упрощается добавление новых типов, например, драйверов

// Минусы
// 1. Может получиться большая иерархия классов с фабриками или большое количество функций, реализующих их
