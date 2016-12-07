package jwt

import (
	"testing"
	"github.com/smhouse/pi/db"
)

func init() {
	db.OpenDatabase("../pi.db")
}

func TestCreateToken(t *testing.T) {

	secret, err := db.GetSecret()
	if err != nil {
		t.Error(err)
	}

	tokenUser := db.User_t{
		Name:		"foo",
		Email:		"foo@gmail.com",
		Password:	"123456",
	}

	token, err := CreateToken(&tokenUser, *secret)
	if err != nil {
		t.Error(err.Error())
		return
	}

	validUser, err := CheckToken(*token, *secret)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if validUser.Name != tokenUser.Name {
		t.Error("User not equals")
		t.Error(validUser.Name)
		t.Error(tokenUser.Name)
	}
}

func TestCheckToken(t *testing.T) {
	secret, err := db.GetSecret()
	if err != nil {
		t.Error(err)
	}

	_, err = CheckToken("fff", *secret)
	if err == nil {
		t.Error("No error on bad token")
		return
	}
}

func TestCheckTokenWrongSign(t *testing.T) {
	secret, err := db.GetSecret()
	if err != nil {
		t.Error(err)
	}

	badToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6Ijk0NTNjNGZhMjg4ODM3MzQxYTBhYWEwYWRlZDE4YjllZGFkZDc1MGQxYmFiYjAwYjgzMWM5OTY5YWM4ZWFlNGViMzIwMWVkZWVkM2M2YmY2In0.eyJhdWQiOiI1MSIsImp0aSI6Ijk0NTNjNGZhMjg4ODM3MzQxYTBhYWEwYWRlZDE4YjllZGFkZDc1MGQxYmFiYjAwYjgzMWM5OTY5YWM4ZWFlNGViMzIwMWVkZWVkM2M2YmY2IiwiaWF0IjoxNDc0MTQ4NDY0LCJuYmYiOjE0NzQxNDg1MjQsImV4cCI6MTQ3NDE1NTY2NCwic3ViIjoiNTciLCJzY29wZXMiOlsiXCJcXFwiYXV0aFxcXCJcIiIsIlwiXFxcImF1dGhcXFwiXCIiXX0.hbVww1lc6lxlyj1HnSLUJUwbpYKUxnHmMdQig-DPyCBF6rG0uiiMHm08Qouc-KHmvZYKOsPgfiXqQMTfRI1jd2R78j2c60YG3voncBu2PxkYvxWLbIeCA6przrmTXWW5kDchIVf1uQuGgfrAis7rqsd0p51CMrxfvheEiAv2HU8"

	_, err = CheckToken(badToken, *secret)
	if err == nil {
		t.Error("No error on bad token")
		return
	}
}