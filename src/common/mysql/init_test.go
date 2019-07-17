package mysql

import (
	"testing"
)

func Test(t *testing.T) {
	tt := &TestTable{Name: "test"}
	id, err := tt.Insert()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id, err)

	tt.Name = "hhh"
	id, err = tt.Update("name")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id, err)
}
