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
)

func init() {
	err := db.OpenDatabase("../pi.db")
	if err != nil {
		log.Fatalln(err)
	}
}

func TestCreateUser(t *testing.T) {
	r := gin.New()
	r.POST("/", CreateUser)

	user_b, _ := json.Marshal(db.User_t{
		Name:		"foo",
		Email:		"foo@bar.com",
		Password:	"bar",
	})
	log.Println(string(user_b))
	req, _ := http.NewRequest("POST", "/", bytes.NewReader(user_b))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Error(w.Body.String())
	}
}