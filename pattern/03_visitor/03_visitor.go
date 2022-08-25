package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
[1] https://medium.com/@felipedutratine/visitor-design-pattern-in-golang-3c142a12945a
[2] https://golangbyexample.com/visitor-design-pattern-go/
Поведенческий паттерн который позволяет добавить поведения структуре НЕ имзеняя ее структуру.
Способ разделения алгоритма от объекта\структуры над которой он оперирует.

1) Когда нет возможности\желания изменять код сторонней библиотеки. Хочется расширить.
2)

	Использование:
	1) когда нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов (дерево)
	2) когда не нужно засорять классы несвязанными операциями
	3) когда новое поведение имеет смысл только для некоторых классов из существующей иерархии
	+:
	1) упрощает добавление операций, работающих со сложными структурами объектов
	2) объединение родственных операций в одном классе
	3) посетитель может накапливать состояние при обходе структуры элементов
	-:
	1) паттерн не оправдан, если иерархия элементов часто меняется
	2) может привести к нарушению инкапсуляции элементов
*/

// даннный интерфейс имплементируем структурами Актер, Сценарист
// [3] - добавляем метод аксепт для допуска визитора
// через accept удается расширить через Визитора интерфейс еще двумя методами
type FilmCrewMember interface {
	GetBiographyNote()
	Accept(Visitor)
}

// [0] - создаем структуру и методы которые будем расширять
type Actor struct {
	Name    string
	Surname string
	Country string
	Fee     float32
	Age     int
}

func (a Actor) GetBiographyNote() {
	fmt.Printf("%s %s was born in  %s. He is %d yaers old. He's day rate is %f  USD\n", a.Name, a.Surname, a.Country, a.Age, a.Fee)
}

type Screenwriter struct {
	Name    string
	Surname string
	Country string
	Fee     float32
	Age     int
}

func (s Screenwriter) GetBiographyNote() {
	fmt.Printf("%s %s was born in  %s. He is %d yaers old. He's day rate is %f USD\n", s.Name, s.Surname, s.Country, s.Age, s.Fee)
}

// [1] - чтобы расширить структуры необходимо создать интрефес Visitor
type Visitor interface {
	VisitActor(a Actor)
	VisitScreenwriter(s Screenwriter)
}

// [2] - каждый дополнительный метод для расширения создаем в отдельную структуру
type CalculateHonorarium struct {
	// считаем сумму ганоррар на основе сьемочнх дней
	filmingDays int
}

func (c CalculateHonorarium) VisitActor(a Actor) {
	fmt.Printf("%s %s income is = %f\n", a.Name, a.Surname, a.Fee*float32(c.filmingDays))
}

func (c CalculateHonorarium) VisitScreenwriter(s Screenwriter) {
	fmt.Printf("%s %s income is = %f\n", s.Name, s.Surname, s.Fee*float32(c.filmingDays))
}

// [4] - реализуем аксепт для каждой структуру . Необходимо для полной реализации интерфейса

func (a Actor) Accept(v Visitor) {
	v.VisitActor(a)
}

func (s Screenwriter) Accept(v Visitor) {
	v.VisitScreenwriter(s)
}

//[5] - добавляем еще один спопосб поведения
type Filmography struct {
	films []string
}

func (f Filmography) VisitActor(a Actor) {
	fmt.Printf("%s %s Films are = %v\n", a.Name, a.Surname, f.films)
}

func (f Filmography) VisitScreenwriter(s Screenwriter) {
	fmt.Printf("%s %s Films are = %v\n", s.Name, s.Surname, f.films)
}

func main() {
	a := Actor{"Leonardo", "DiCaprio", "California, USA", 100000, 47}
	s := Screenwriter{"Quentin", "Tarantino", "Tennessi, USA", 102000, 59}
	a.GetBiographyNote()
	s.GetBiographyNote()

	a.Accept(CalculateHonorarium{78})

	s.Accept(CalculateHonorarium{10})

	fl := Filmography{
		films: []string{"The Wolf of Wall Street", "Django Unchained"},
	}
	a.Accept(fl)

	s.Accept(Filmography{
		films: []string{"Pulp Fiction", "Django Unchained"},
	})
}
