package db

import "testing"

func TestDevice_t_Create(t *testing.T) {
	d := Device_t{
		Name:		"dev_1",
		Password:	"123456",
		Description:	"sadasdas",
	}
	err := d.Create()

	if err != nil {
		t.Error(err)
	}

	df := Device_t{ID: d.ID}
	err = df.Find()
	if err != nil {
		t.Error(err)
	}

	err = d.Delete()
	if err != nil {
		t.Error(err)
	}
}
