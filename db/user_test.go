package db

import (
	"testing"
	"log"
)

func init() {
	err := OpenDatabase("../pi.db")
	if err != nil {
		log.Fatalln(err)
	}
}

func TestUser_t_Create(t *testing.T) {
	u := User_t{
		Name:		"test",
		Email:		"test@test.com",
		Password:	"123456",
	}
	err := u.Create()
	if err != nil {
		t.Error(err)
	}

	err = u.Delete()
	if err != nil {
		t.Error(err)
	}
}

func TestUser_t_Find(t *testing.T) {
	u := User_t{
		Name:		"test",
		Email:		"test@test.com",
		Password:	"123456",
	}
	err := u.Create()
	if err != nil {
		t.Error(err)
	}

	findUser := User_t{Name: u.Name}
	err = findUser.Find()
	if err != nil {
		t.Error(err)
	}

	if findUser.Email != u.Email {
		t.Error("Users not equals")
	}
}