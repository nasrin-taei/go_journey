package main

import "fmt"

type Account struct {
	id   int //Shadowed field
	name string
}

func (a Account) String() string { //shadowed function
	return fmt.Sprintf("Account with id %d belongs to %s", a.id, a.name)
}

func (a *Account) changeName(name string) {
	a.name = name
}

type ManagerAccount struct {
	Account
	id int
}

func (m ManagerAccount) String() string {
	m.changeName(m.name + ".Jr")
	return fmt.Sprintf("Manager Account with id %d belongs to %s", m.id, m.name)
}

func main() {
	mgr := ManagerAccount{Account: Account{10, "Jack"}, id: 15}
	fmt.Println(mgr)
	mgr.changeName("Jack Dorsey")
	fmt.Println(mgr)
}
