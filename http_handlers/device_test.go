package http_handlers

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/smhouse/pi/db"
	"github.com/smhouse/pi/jwt"
	"net/http"
	"encoding/json"
	"bytes"
	"net/http/httptest"
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
}