package env_reader

import (
	"fmt"
	"reflect"
)

type ValueReader = func(string) string

func Read(r ValueReader, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer, but got %v", val.Kind())
	}

	t := val.Elem().Type()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("expected pointer to struct, but got %v", t.Kind())
	}

	//TODO add reading nested structures
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if tag := f.Tag.Get("env"); len(tag) > 0 {
			if s := r(tag); len(s) > 0 {
				val.Elem().Field(i).SetString(s)
			}
		}
	}

	return nil
}
