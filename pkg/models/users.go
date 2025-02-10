package users

import "fmt"

type Users struct {
	Name string
	DOB  int
}

func (u *Users) String() {
	fmt.Println("Hello user, I am version 1")
}
