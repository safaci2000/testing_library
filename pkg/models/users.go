package models

import "fmt"

type Users struct {
	Name string
}

func (u *Users) String() string {
	return fmt.Sprintf("Hello %s, I am version 2.", u.Name)
}
