package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Использование:
	1) когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего
	состояния, причем типов состояния много, и их код часто меняется
	2) когда код класса содержит множество больших, похожих друг на друга, условных операторов,
	которые выбирают поведения в зависимости от текущих значений полей класса
	3) когда сознательно используют табличную машину состояний, построенную на условных операторах,
	но вынуждены мириться с дублированием кода для похожих состояний и переходов
	+:
	1) избавляет от множества больших условных операторов машины состояний
	2) концентрирует в одном месте код, связанный с определенным состоянием
	3) упрощает код контекста
	-:
	1) может неоправданно усложнить код, если состояний мало и они редко меняются
*/

/*
[1] https://golangbyexample.com/state-design-pattern-go/
[2] https://venilnoronha.io/a-simple-state-machine-framework-in-go

Паттерн состояния хорошо применять, когда есть множество одинаковых действий, которые
В зависимости от множества разных состояний делают разные вещи
Паттерн состояния позволяет разделить состояния на классы и уже через них выполнять какие-либо действия
Как раз в зависимости от состояния
Это очень похоже предыдущий паттерн, но там стратегию выбирает сам клиент, а здесь "машина" с состоянием
Сама решает, что делает в зависимости от состояния, которое указано
*/

// контролируем агрегатное состояние  ртути

//[0] - интерфеес состояния
type IMercuryState interface {
	Freeze()
	Heat()
}

//[1] - реализуем конкретное состояние в виде структуры и методов
//[1.1] - твердое
type SolidMercuryState struct {
	mercury *Mercury
}

func (ms *SolidMercuryState) Freeze() {
	fmt.Println("[SolidMercuryState]:: Freezing Mercury ...")
}

func (ms *SolidMercuryState) Heat() {
	fmt.Println("[SolidMercuryState]:: Heating Mercury to liquid state...")
	ms.mercury.SetState(&ms.mercury.liquidMercury)
}

//[1.2]
type LiquidMercuryState struct {
	mercury *Mercury
}

func (ls *LiquidMercuryState) Freeze() {
	fmt.Println("[LiquidMercuryState]:: Freezing mercury to crystal...")
	ls.mercury.SetState(&ls.mercury.freezedMercury)
}
func (ls *LiquidMercuryState) Heat() {
	fmt.Println("[LiquidMercuryState]:: Evaporating liquid mercury to steam...")
	ls.mercury.SetState(&ls.mercury.gaseousMercury)
}

//[1.3]
type GaseousMercuryState struct {
	mercury *Mercury
}

func (gs *GaseousMercuryState) Freeze() {
	fmt.Println("[GaseousMercuryState]:: Condensing Mercury Steam to Liquid Mercury...")
	gs.mercury.SetState(&gs.mercury.liquidMercury)
}
func (gs *GaseousMercuryState) Heat() {
	fmt.Println("[GaseousMercuryState]:: Making Mercury steam hotter...")
}

//==========================================

/*
[2] - Context structure, context struct initialization
implement StateInterface methods + implement setState() method
*/

type Mercury struct {
	freezedMercury SolidMercuryState
	liquidMercury  LiquidMercuryState
	gaseousMercury GaseousMercuryState

	currentState IMercuryState
}

func NewMercury() *Mercury {
	mercuryMolecule := &Mercury{}
	solid := SolidMercuryState{}
	liquid := LiquidMercuryState{}
	gas := GaseousMercuryState{}

	mercuryMolecule.freezedMercury = solid
	mercuryMolecule.liquidMercury = liquid
	mercuryMolecule.gaseousMercury = gas

	return mercuryMolecule
}

func (m *Mercury) SetState(state IMercuryState) {
	m.currentState = state
}

func (m *Mercury) Freeze() {
	m.currentState.Freeze()
}
func (m *Mercury) Heat() {
	m.currentState.Heat()
}

func main() {
	mercury := NewMercury()
	mercury.Heat()
	mercury.Heat()
	mercury.Freeze()
	mercury.Freeze()
	mercury.Freeze()

}
