package db

import "testing"

func TestGetSecret(t *testing.T) {
	_, err := GetSecret()
	if err != nil {
		t.Error(err)
	}
}
