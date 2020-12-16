package main

import (
	"container/list"
	"fmt"
	"time"
)

func simple() {

	lst := list.New()

	e4 := lst.PushBack(4)
	e1 := lst.PushFront(1)

	lst.InsertBefore(3, e4)
	lst.InsertAfter(2, e1)

	fmt.Printf("length: %d\n", lst.Len())

	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

var id int

type User struct {
	name string
	ttl  int64
	id   int
}

func NewUser(name string) *User {
	user := new(User)
	user.id = id
	user.name = name
	user.ttl = time.Now().UnixNano()

	time.Sleep(time.Millisecond)

	id++

	return user
}

func complex() {
	id = 100
	// list of user objects
	users := list.New()

	// add some users
	users.PushFront(NewUser("tom"))

	// get a reference to a user
	dick := NewUser("dick")
	users.PushFront(dick)

	// add another user
	users.PushFront(NewUser("jane"))

	// show the users in order of insertion
	for e := users.Front(); e != nil; e = e.Next() {
		user, ok := e.Value.(*User)

		fmt.Printf("%v id: %d name: %s ttl: %d\n", ok, user.id, user.name, user.ttl)
	}

	// remove a specific element
	for e := users.Front(); e != nil; e = e.Next() {
		user, ok := e.Value.(*User)
		if ok && user.name == dick.name {
			fmt.Println("remove dick...")
			users.Remove(e)
			break
		}
	}

	// show the modified list
	fmt.Println("post dick...")
	for e := users.Front(); e != nil; e = e.Next() {
		user, ok := e.Value.(*User)

		fmt.Printf("%v id: %d name: %s ttl: %d\n", ok, user.id, user.name, user.ttl)
	}

	if user, ok := users.Back().Value.(*User); ok {
		fmt.Println("back of the list:")
		fmt.Printf("%v id: %d name: %s ttl: %d\n", ok, user.id, user.name, user.ttl)
	}
}

func main() {
	// simple()
	complex()
}
