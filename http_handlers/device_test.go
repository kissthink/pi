package http_handlers

import (
	"testing"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/smhouse/pi/db"
	"github.com/smhouse/pi/jwt"
	"net/http"
	"encoding/json"
	"bytes"
	"net/http/httptest"
	"fmt"
)

func TestCreateDevice(t *testing.T) {

	u := db.User_t{
		Name: "test",
		Email: "test@test.com",
		Password: "123456",
	}
	err := u.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer u.Delete()

	token, err := jwt.CreateToken(&u)
	if err != nil {
		t.Error(err)
		return
	}

	d := db.Device_t{
		Name: "dev1",
		Password: "123456",
	}

	device_b, _ := json.Marshal(d)

	r := gin.New()
	r.Use(AuthJWT())
	r.POST("/", CreateDevice)

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(device_b))
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

	newDevice := db.Device_t{}
	err = json.Unmarshal(w.Body.Bytes(), &newDevice)
	if err != nil {
		t.Error(err)
		return
	}

	err = newDevice.Delete()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteDevice(t *testing.T) {
	u := db.User_t{
		Name: "test",
		Email: "test@test.com",
		Password: "123456",
	}
	err := u.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer u.Delete()

	d := db.Device_t{
		Name: "foo",
		Password: "bar",
		UserName: u.Name,
	}
	err = d.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer d.Delete()

	token, err := jwt.CreateToken(&u)
	if err != nil {
		t.Error(err)
		return
	}

	r := gin.New()
	r.Use(AuthJWT())
	r.DELETE("/:id", DeleteDevice)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/%d", d.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Token", *token)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)


	if status := w.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
		t.Error(w.Body.String())
		return
	}
}

func TestUpdateDevice(t *testing.T) {
	u := db.User_t{
		Name: "test",
		Email: "test@test.com",
		Password: "123456",
	}
	err := u.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer u.Delete()

	d := db.Device_t{
		Name: "foo",
		Password: "bar",
		UserName: u.Name,
	}
	err = d.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer d.Delete()

	token, err := jwt.CreateToken(&u)
	if err != nil {
		t.Error(err)
		return
	}

	r := gin.New()
	r.Use(AuthJWT())
	r.PUT("/:id", UpdateDevice)

	d.Password = "123456"
	data, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
	}

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/%d", d.ID), bytes.NewReader(data))
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

func TestListDevices(t *testing.T) {
	u := db.User_t{
		Name: "test",
		Email: "test@test.com",
		Password: "123456",
	}
	err := u.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer u.Delete()

	d := db.Device_t{
		Name: "foo",
		Password: "bar",
		UserName: u.Name,
	}
	err = d.Create()
	if err != nil {
		t.Error(err)
		return
	}

	defer d.Delete()

	token, err := jwt.CreateToken(&u)
	if err != nil {
		t.Error(err)
		return
	}

	r := gin.New()
	r.Use(AuthJWT())
	r.GET("/", ListDevices)

	req, _ := http.NewRequest("GET", "/", nil)
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

	devices := make([]db.Device_t, 0)
	err = json.Unmarshal(w.Body.Bytes(), &devices)
	if err != nil {
		t.Error(err)
	}

	if len(devices) != 1 {
		t.Error("Wrong devices length")
		fmt.Printf("%v\n", devices)
	}
}