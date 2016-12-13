package http_handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/smhouse/pi/db"
	"encoding/json"
	"bytes"
	"fmt"
	"strings"
	"github.com/smhouse/pi/jwt"
)

func init() {
	err := db.OpenDatabase("../pi.db")
	if err != nil {
		log.Fatalln(err)
	}
	gin.SetMode("test")
	db.GetSecret()
}

func TestCreateUser(t *testing.T) {
	r := gin.New()
	r.POST("/", CreateUser)

	form := user_create_form{
		Name:		"foo123",
		Email:		"foo123@bar.com",
		Password:	"bar123",
	}

	user_b, _ := json.Marshal(form)

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(user_b))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Error(w.Body.String())
		return
	}

	req, _ = http.NewRequest("POST", "/", bytes.NewReader(user_b))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		return
	}

	u := db.User_t{Name: form.Name}
	err := u.Delete()
	if err != nil {
		t.Error(err)
		return
	}

	req, _ = http.NewRequest("POST", "/", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		return
	}
}

func TestGetUser(t *testing.T) {
	u := db.User_t{
		Name:		"foo",
		Email:		"foo@gmail.com",
		Password:	"123456",
	}

	err := u.Create()
	if err != nil {
		t.Error(err)
	}

	r := gin.New()
	r.GET("/:name", GetUser)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", u.Name), nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Error(w.Body.String())
		return
	}

	if err := u.Delete(); err != nil {
		t.Error(err)
		return
	}

	req, _ = http.NewRequest("GET", fmt.Sprintf("/%s", u.Name), nil)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		return
	}
}

func TestLoginUser(t *testing.T) {
	u := db.User_t{
		Name:		"foo",
		Email:		"foo@gmail.com",
		Password:	"123456",
	}

	err := u.Create()
	if err != nil {
		t.Error(err)
	}

	r := gin.New()
	r.POST("/", LoginUser)

	form := user_login_form{
		Name:		"foo",
		Password:	"123456",
	}

	user_b, _ := json.Marshal(form)

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(user_b))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Error(w.Body.String())
		return
	}

	if err := u.Delete(); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	u := db.User_t{
		Name:		"foo",
		Email:		"foo@gmail.com",
		Password:	"123456",
	}

	err := u.Create()
	if err != nil {
		t.Error(err)
	}

	defer u.Delete()

	token, err := jwt.CreateToken(&u)
	if err != nil {
		t.Error(err)
		return
	}

	r := gin.New()
	r.PUT("/", AuthJWT(), ListDevices)

	form := user_update_form{
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
		NewPassword: "123456789",
	}

	data, err := json.Marshal(form)
	if err != nil {
		t.Error(err)
		return
	}

	req, _ := http.NewRequest("PUT", "/", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", *token)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Error(w.Body.String())
		return
	}
}