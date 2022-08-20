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
[3] https://blog.ralch.com/articles/design-patterns/golang-builder/
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

	GetServerStructure() *Server // возвращает полную структуру сервера
}

/*
Director implements a manager.
Директор ( тут заказчик), определяет ПОРЯДОК в котором ВЕНДОР будет собирать сервер
*/
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in what order. Их может быть НЕСКОЛЬКО
func (d *Director) ConfigureCPUServer(builder Builder, cpu string, cores, sk, ram int, os string) *Server {
	builder.ChooseCPU(cpu)
	builder.ChooseCores(cores)
	builder.ChooseSockets(sk)
	builder.ChooseRAM(ram)
	builder.ChooseOS(os)

	return builder.GetServerStructure()
}

// Второй метод. Конфигурируем с ГПУ
func (d *Director) ConfigureGPUServer(builder Builder, cpu, gpu string, gpunumber, cores int, os string) *Server {
	builder.ChooseCPU(cpu)
	builder.ChooseCores(cores)
	builder.ChooseGPU(gpu)
	builder.ChooseGPUNumber(gpunumber)
	builder.ChooseOS(os)

	return builder.GetServerStructure()
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
func (lb *LenovoBuilder) ChooseCores(coreCount int) {
	lb.configuration.Cores = coreCount
	fmt.Printf("Lenovo has chosen CPU with %d cores\n", lb.configuration.Cores)
}

// имплементировать КАЖДЫЙ МЕТОД
func (lb *LenovoBuilder) ChooseSockets(socketsNum int) {}
func (lb *LenovoBuilder) ChooseRAM(gbs int)            {}
func (lb *LenovoBuilder) ChooseGPU(gpuModel string) {
	lb.configuration.GPU = gpuModel
	fmt.Printf("Lenovo has chosen GPU model : %s\n", lb.configuration.GPU)
}
func (lb *LenovoBuilder) ChooseGPUNumber(num int) {}
func (lb *LenovoBuilder) ChooseHDDMemory(gbs int) {}
func (lb *LenovoBuilder) ChooseOS(os string)      {}
func (lb *LenovoBuilder) ChooseGoal(goal string)  {}

func (lb *LenovoBuilder) GetServerStructure() *Server {
	fmt.Printf("Lenovo proposed configuration  : %v \n", lb.configuration)
	return &lb.configuration
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
func (hp *HPBuilder) ChooseCores(coreCount int) {
	hp.configuration.Cores = coreCount
	fmt.Printf("Lenovo has chosen CPU with %d cores\n", hp.configuration.Cores)
}

// имплементировать КАЖДЫЙ МЕТОД
func (hp *HPBuilder) ChooseSockets(socketsNum int) {}
func (hp *HPBuilder) ChooseRAM(gbs int)            {}
func (hp *HPBuilder) ChooseGPU(cpuModel string)    {}
func (hp *HPBuilder) ChooseGPUNumber(num int)      {}
func (hp *HPBuilder) ChooseHDDMemory(gbs int)      {}
func (hp *HPBuilder) ChooseOS(os string)           {}
func (hp *HPBuilder) ChooseGoal(goal string)       {}

func (hp *HPBuilder) GetServerStructure() *Server {
	fmt.Printf("HP proposed configuration  : %v \n", hp.configuration)
	return &hp.configuration
}

// ==================================================================================

// 3. Реализация двух  билдеров  : GPU Server, CPU Server
// ==================================================================================

func main() {

	lenovoBuilder := &LenovoBuilder{}
	lenovoBuilder.ChooseCPU("Intel CLX8260L")
	lenovoBuilder.ChooseCores(48)
	lenovoBuilder.ChooseSockets(2)
	lenovoBuilder.ChooseRAM(192)

	resultLenovo := lenovoBuilder.GetServerStructure()
	fmt.Printf("%v\n", resultLenovo)

	hpBuilder := &HPBuilder{}
	hpBuilder.ChooseCPU("AMD EPYC 7763")
	hpBuilder.ChooseCores(64)
	hpBuilder.ChooseSockets(2)
	hpBuilder.ChooseRAM(192)
	resultHP := hpBuilder.GetServerStructure()
	fmt.Printf("%v\n", resultHP)

	fmt.Println("==========USE DIRECTOR===================")
	// через Директора !
	// Директор нужен каждому билдеру?
	dir := &Director{}
	res := dir.ConfigureCPUServer(&HPBuilder{}, "IBM Power S814", 6, 2, 16, "IBM Solaris")
	fmt.Println(res)

	res2 := dir.ConfigureGPUServer(&LenovoBuilder{}, "Intel Xeon v3", "Nvidia V100", 2, 12, "RedHat 9.2")
	fmt.Println(res2)

	// через лонг билдер
	// HomoServer := NewServerBuilder{}
	// HomoServer.HomogenousConstructor().ChooseCPU("").ChooseCores(1)
	// s := HomoServer.Build()
	// fmt.Println(s)

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
