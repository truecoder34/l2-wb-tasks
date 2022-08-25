package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Поведенческий паттерн - позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи
*/

/*
[1] https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/
[2] https://russianblogs.com/article/24892334053/

Есть структура данных которая приходит от пользователя, с фронта, DTO.
- часть данных передаем в логер,
- часть данных пишем в базу,
- часть данных передаем в меседж брокер
*/

// [0] - абстрактный интерфейс обработчика - Handler
type handlerDTO interface {
	execute(*MessageDTO)
	setNext(handlerDTO)
}

//========================================================

//[1] - msg dto that we receive and will handle
type MessageDTO struct {
	id                string
	userNickSender    string
	userEmailSender   string
	message           string
	userNickReceiver  string
	userEmailReceiver string
}

//========================================================

//[2] - create logger structure and implement handler interface
type Logger struct {
	nextHandler   handlerDTO
	id            string
	senderEmail   string
	receiverEmail string
}

func (l *Logger) execute(m *MessageDTO) {
	l.id = m.id
	l.senderEmail = m.userEmailSender
	l.receiverEmail = m.userEmailReceiver

	fmt.Printf("[Logger]:: Pushing message %s from %s to %s \n", l.id, l.senderEmail, l.receiverEmail)

	//l.nextHandler.execute(m)
}

func (l *Logger) setNext(nextH handlerDTO) {
	l.nextHandler = nextH
}

//========================================================

//[3] - create DB structure and implement handler interface
type DataBase struct {
	nextHandler   handlerDTO
	nickSender    string
	emailSender   string
	nickReceiver  string
	emailReceiver string
}

func (db *DataBase) execute(m *MessageDTO) {
	db.emailReceiver = m.userEmailReceiver
	db.emailSender = m.userEmailSender
	db.nickReceiver = m.userNickReceiver
	db.nickSender = m.userNickReceiver

	fmt.Printf("[DataBase Controller]:: Data write to DB: sender email - %s, receiver email - %s, sender nick - %s, receiver nick - %s\n", db.emailSender, db.emailReceiver, db.nickSender, db.nickReceiver)

	db.nextHandler.execute(m)
}

func (db *DataBase) setNext(nextH handlerDTO) {
	db.nextHandler = nextH
}

//========================================================

//[4] - create MsgBrocker structure and implement handler interface
type MsgBroker struct {
	nextHandler handlerDTO
	id          string
	msg         string
}

func (mb *MsgBroker) execute(m *MessageDTO) {
	mb.id = m.id
	mb.msg = m.message
	fmt.Printf("[Message Broker]:: Send hashed message to storage . ID : %s ; Body : %s t\n", mb.id, mb.msg)

	mb.nextHandler.execute(m)
}

func (mb *MsgBroker) setNext(nextH handlerDTO) {
	mb.nextHandler = nextH
}

//========================================================

func main() {
	lg := &Logger{}
	db := &DataBase{}
	mb := &MsgBroker{}

	msgDTO := &MessageDTO{
		id:                "1111",
		userNickSender:    "user1",
		userEmailSender:   "user1@mail.com",
		message:           "MSG MSG MSG HELLO",
		userNickReceiver:  "user2",
		userEmailReceiver: "user2@mail.com",
	}

	db.setNext(lg)
	mb.setNext(db)

	mb.execute(msgDTO)

}

/*
	Использование:
	1) когда программа должна обрабатывать разнообразные запросы несколькими способами,
	но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся
	2) когда важно, чтобы обработчики выполнялись один за другим в строгом порядке
	3) когда набор объектов, способных обработать запрос, должен задаваться динамически
	+:
	1) уменьшает зависимость между клиентом и обработчиками
	2) реализует принцип единственной обязаности
	3) реализует принцип открытости/закрытости
	-:
	1) запрос может остаться никем не обработанным
*/
