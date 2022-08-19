package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Examples
[1] https://devcharmander.medium.com/design-patterns-in-golang-the-builder-dac468a71194
[2] https://github.com/AlexanderGrom/go-patterns/blob/master/Creational/Builder/builder.go
*/

/*
КЛАСС КОНСТРУИРУЕМОГО ПРОДУКТА :
Собираем сервер для хостинга или вычисления. Список характеристик урезан.
Их много больше, это и есть сложная компонента которую будем строить с помощью паттерна
*/
type Server struct {
	CPU       string
	Cores     int
	Sockets   int
	RAM       int
	GPU       string
	GPUNumber int
	HDD       int
	OS        string
	Goal      string
}

/*
Сложный конструктор для данной структуры.НЕ ВСЕ элементы будут задейстованы
так делать не удобно : Сервер может быть без ГПУ. Часть с ГПУ можно вынести отдельно
*/
func InitServer(cpu string, cores, sockets int, ram int, gpu string, gpunumber int, hdd int, os, goal string) Server {
	return Server{
		CPU:       cpu,
		Cores:     cores,
		Sockets:   sockets,
		RAM:       ram,
		GPU:       gpu,
		GPUNumber: gpunumber,
		HDD:       hdd,
		OS:        os,
		Goal:      goal,
	}
}

// ==================================================================================

/*
Общий интерфейс для строителей!
Применяем паттерн строитель BUILDER  для данной структуры
интерфейс для билдера. Методы для выбора характеристик сервера
*/
type Builder interface {
	ChooseCPU(cpuModel string)
	ChooseSockets(socketsNum int)
	ChooseCores(coreCount int)
	ChooseRAM(gbs int)
	ChooseGPU(cpuModel string)
	ChooseGPUNumber(num int)
	ChooseHDDMemory(gbs int)
	ChooseOS(os string)
	ChooseGoal(goal string) // storage, computing

	GetServerStructure() Server
}

/*
Director implements a manager.
Директор ( тут заказчик), определяет ПОРЯДОК в котором ВЕНДОР будет собирать сервер
*/
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in what order. Их может быть НЕСКОЛЬКО
func (d *Director) ConfigureCPUServer(cpu string, cores int, sockets int, ram int, os string) {
	d.builder.ChooseCPU(cpu)
	d.builder.ChooseCores(cores)
	d.builder.ChooseSockets(sockets)
	d.builder.ChooseRAM(ram)
	d.builder.ChooseOS(os)
}

// Второй метод. Конфигурируем с ГПУ
func (d *Director) ConfigureGPUServer() {

}

// ==================================================================================

// ConcreteBuilder implements Builder interface .В данном случае ВЕНДОР - Lenovo, HP, HPE - Те кто будут собирать сервер для нас
// ==================================================================================
type LenovoBuilder struct {
	configuration Server
}

// Конкретный билдер имплементирует каждый метод интерфейса
func (lb *LenovoBuilder) ChooseCPU(model string) {
	lb.configuration.CPU = model
	fmt.Printf("Lenovo has chosen CPU model : %s\n", lb.configuration.CPU)
}

// имплементировать КАЖДЫЙ МЕТОД
func (lb *LenovoBuilder) ChooseSockets(socketsNum int) {}
func (lb *LenovoBuilder) ChooseCores(coreCount int)    {}
func (lb *LenovoBuilder) ChooseRAM(gbs int)            {}
func (lb *LenovoBuilder) ChooseGPU(cpuModel string)    {}
func (lb *LenovoBuilder) ChooseGPUNumber(num int)      {}
func (lb *LenovoBuilder) ChooseHDDMemory(gbs int)      {}
func (lb *LenovoBuilder) ChooseOS(os string)           {}
func (lb *LenovoBuilder) ChooseGoal(goal string)       {}

func (lb *LenovoBuilder) GetServerStructure() Server {
	fmt.Printf("Lenovo proposed configuration  : %v \n", lb.configuration)
	return lb.configuration
}

// ==================================================================================

// 2. Реализация еще одного билдера (Вендор) - HP
// ==================================================================================
type HPBuilder struct {
	configuration Server
}

func (hp *HPBuilder) ChooseCPU(model string) {
	hp.configuration.CPU = model
	fmt.Printf("HP has chosen CPU model : %s\n", hp.configuration.CPU)
}

// имплементировать КАЖДЫЙ МЕТОД
func (hp *HPBuilder) ChooseSockets(socketsNum int) {}
func (hp *HPBuilder) ChooseCores(coreCount int)    {}
func (hp *HPBuilder) ChooseRAM(gbs int)            {}
func (hp *HPBuilder) ChooseGPU(cpuModel string)    {}
func (hp *HPBuilder) ChooseGPUNumber(num int)      {}
func (hp *HPBuilder) ChooseHDDMemory(gbs int)      {}
func (hp *HPBuilder) ChooseOS(os string)           {}
func (hp *HPBuilder) ChooseGoal(goal string)       {}

func (hp *HPBuilder) GetServerStructure() Server {
	fmt.Printf("HP proposed configuration  : %v \n", hp.configuration)
	return hp.configuration
}

func NewHP() Builder {
	return &HPBuilder{}
}

func main() {

	LenovoBuilder := &LenovoBuilder{}
	LenovoBuilder.ChooseCPU("Intel CLX8260L")
	LenovoBuilder.ChooseCores(48)
	LenovoBuilder.ChooseSockets(2)
	LenovoBuilder.ChooseRAM(192)
	resultLenovo := LenovoBuilder.GetServerStructure()
	fmt.Printf("%v\n", resultLenovo)

	HPBuilder := &HPBuilder{}
	HPBuilder.ChooseCPU("AMD EPYC 7763")
	HPBuilder.ChooseCores(64)
	HPBuilder.ChooseSockets(2)
	HPBuilder.ChooseRAM(192)
	resultHP := HPBuilder.GetServerStructure()
	fmt.Printf("%v\n", resultHP)

	// через Директора !
	// Директор нужен каждому билдеру?
	// dir := &Director{}
	// dir.ConfigureCPUServer("IBM Power S814", 6, 2, 16, "IBM Solaris")

	// через лонг билдер
	//HPBuilder2 = &LenovoBuilder{}
	//HPBuilder2.ChooseCPU("AMD EPYC 7763").ChooseCores(64).ChooseSockets(2)
	//HPBuilder2 := NewHP()
	//HPBuilder2.ChooseCPU("AMD EPYC 7763").ChooseCores(64).ChooseSockets(2)
}

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Строитель позволяет не использовать длинные конструкторы.

Если потребуется объединить шаги создания создается объект директор, определяющий
порядок шагов, так процесс создания будет полностью скрыт от клиента
(директор - не является обязательным в реализации паттерна)

Плюсы:
- Создает объекты пошагово (если это требуется, например, деревья)
- Позволяет создавать несколько представлений одного объекта, переиспользуя код

Минусы:
- Усложняет код, требуется написание дополнительных классов, может усложняться внедрнее зависимостей
*/
