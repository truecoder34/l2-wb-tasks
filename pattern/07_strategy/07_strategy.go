package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
[1] https://golangbyexample.com/strategy-design-pattern-golang/
[2] https://golang-blog.blogspot.com/2019/10/strategy-pattern.html
*/

/*
Поведенческий паттерн
Через общий интерфейс получать результат
С помощью разных алгоритмов или методов, но через один общий интерфейс
Использование:
	1) когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта
	2) когда у вас есть множество похожих классов, отличающихся только некоторым поведением
	3) когда не нужно обнажать детали реализации алгоритмов для других классов
	4) когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
	каждая ветка такого оператора представляет собой вариацию алгоритма
+:
	1) горячая замена алгоритмов на лету
	2) изолирует код и данные алгоритмов от остальных классов
	3) уход от наследования к делегированию
	4) реализует принцип открытости/закрытости
-:
	1) усложняет программу за счет дополнительных классов
	2) клиент должен знать, в чем состоит разница между стратегиями, чтобы выбрать подходящую
*/

// Читаем цены на товары в зависимости от скидочной программы
const (
	B_FRIDAY_DISCOUNT = 30
	C_MONDAY_DISCOUNT = 22
	NEW_YEAR_DISCOUNT = 44
	ITEM_COST         = 102400
)

// [0] - Базовый интерфейс стратегии
type iCostStrategy interface {
	calculatePrice(float32) float32
}

// [1] - реализуем интерфейс для различных алгоритмов
type BlackFridayCost struct{}

func (bfc *BlackFridayCost) calculatePrice(amount float32) float32 {
	res := amount - amount/100*B_FRIDAY_DISCOUNT
	fmt.Printf("New cost on BLACK FRIDAY : %f \n", res)
	return res
}

type CyberMondayCost struct{}

func (cmc *CyberMondayCost) calculatePrice(amount float32) float32 {
	res := amount - amount/100*C_MONDAY_DISCOUNT
	fmt.Printf("New cost on CYBER MONDAY : %f \n", res)
	return res
}

type NewYearCost struct{}

func (bfc *NewYearCost) calculatePrice(amount float32) float32 {
	res := amount - amount/100*NEW_YEAR_DISCOUNT
	fmt.Printf("New cost on NEW YEAR : %f \n", res)
	return res
}

// [2] - текущее состояние - текущая стратегия
type ApplicationStratrgy struct {
	CurrentStrategy iCostStrategy
}

func main() {
	app := ApplicationStratrgy{}

	app.CurrentStrategy = &BlackFridayCost{}
	app.CurrentStrategy.calculatePrice(ITEM_COST)

	app.CurrentStrategy = &CyberMondayCost{}
	app.CurrentStrategy.calculatePrice(ITEM_COST)

	app.CurrentStrategy = &NewYearCost{}
	app.CurrentStrategy.calculatePrice(ITEM_COST)
}

// Плюсы
// 1. Изоляция реализации конкретной стратегии
// 2. Возможность менять стратегию "на лету"

// Минусы
// 1. Клиенту (разработчику, не связанному с пакетом стратегии) нужно знать, в чём заключаются стратегии
// 2. Если стратегий будет много и они будут меняться, то код засорится классами
