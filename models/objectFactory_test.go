package model

import (
	"reflect"
	"testing"
)

func TestCreateObjects(t *testing.T) {
	n := 20
	v := createObjects(Event{}, n)
	if len(v) != n {
		t.Errorf("Expected size of %d, got size of %d ", n, len(v))
	}
	for _, a := range v {
		if a == nil {
			t.Errorf("%p is null", a)
		}
	}
}

func TestCreateObject(t *testing.T) {
	objects := []interface{}{Zoom{}, Address{}, Event{}, User{}, Message{}, Guest{}}
	for _, v := range objects {
		t.Run(reflect.TypeOf(v).String(), func(t *testing.T) {
			a := createObject(v)
			if a == nil {
				t.Errorf("%p is null", a)
			}
		})
	}
}

func TestStruct(t *testing.T) {
	id := "28282"
	pass := "ssss"
	z := Zoom{id, pass}
	if z.ZoomId != id {
		t.Errorf("id = %s; want %s", id, z.ZoomId)
	}
	if z.Password != pass {
		t.Errorf("id = %s; want %s", id, z.ZoomId)
	}
}
