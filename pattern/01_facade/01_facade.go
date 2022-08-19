package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

//Account struct
type Account struct {
	id          string
	accountType string
	balance     int
}

//create account
func (ac *Account) createAccount(acType string) *Account {
	rand.Seed(time.Now().UnixNano())
	ac.accountType = acType
	ac.balance = 0
	ac.id = "4276 " + strconv.Itoa(rand.Intn(10000000))
	fmt.Printf("account creation with type = %s\n", ac.accountType)
	return ac
}

//get account by id
func (ac *Account) getAccountById(id string) *Account {
	fmt.Printf("getting account by Id = %s\n", id)
	return ac
}

//add money to account
func (ac *Account) addMoney(id string) *Account {
	ac.balance += 100
	fmt.Printf("account balance increased by 100 ( Id = %s)\n", id)
	return ac
}

//User struct . User can have accounts
type User struct {
	name  string
	email string
	id    int
}

//User create
func (u *User) createUser(n string, e string) *User {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("creating user : %s , %s\n", n, e)
	u.name = n
	u.email = e
	u.id = rand.Intn(10000000)
	return u
}

// Item  to buy
type Item struct {
	name string
	cost int
	id   int
}

//create Item
func (i *Item) create(n string, c int) *Item {
	fmt.Printf("creating Item : %s , cost = %d\n", n, c)
	i.id = rand.Intn(10000000)
	i.cost = c
	i.name = n
	return i
}

//CREATING FACADE FOR THIS CLASSES
type AdminFacade struct {
	user    *User
	account *Account
	item    *Item
}

//Init Facade
func NewAdminFacade() *AdminFacade {
	return &AdminFacade{
		&User{},
		&Account{},
		&Item{},
	}
}

//Create user and account related
func (facade *AdminFacade) createUserAccount(userName string, userEmail string, accountType string) (*User, *Account) {
	var user = facade.user.createUser(userName, userEmail)
	var account = facade.account.createAccount(accountType)
	return user, account
}

// create Item
func (facade *AdminFacade) createItem(name string, cost int) *Item {
	return facade.item.create(name, cost)
}

func main() {
	// 0 - init facade
	facade := NewAdminFacade()
	// 1 - call facade methods
	user, account := facade.createUserAccount("Vlad", "Plotnikov", "User")
	fmt.Println(user.name, user.email)
	fmt.Println(account.accountType)
	// 2 - call facade methods
	item := facade.createItem("Soap", 55)
	fmt.Println(item.name, item.cost)
}
