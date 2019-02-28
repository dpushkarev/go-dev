package main

import (
	"bytes"
	"reflect"
	"encoding/binary"
	"fmt"
	"testing"
)

type User struct {
	ID       int
	Name     string
	LastName string `unpack:"-"`
	Flag     int
}
var data = []byte{
	128, 45, 34, 0,
	9, 0, 0, 0,
	120, 46, 120, 120, 120, 120, 120, 120, 120,
	10, 0, 0, 0,
}
func BrnchUnpackReflect2(b *testing.B){
	for i := 0; i < b.N; i++ {
		u := new(User)
		UnpackReflect2(u, data)
	}
}

func BrnchUnpackReflect22(b *testing.B){
	for i := 0; i < b.N; i++ {
		u := new(User)
		UnpackReflect2(u, data)
	}
}

func UnpackReflect2(u interface{}, data []byte) error {
	r := bytes.NewReader(data)

	val := reflect.ValueOf(u).Elem()

	for i := 0; i < val.NumField(); i++{
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Tag.Get(`unpack`) == `-` {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			var value uint32
			binary.Read(r, binary.LittleEndian, &value)
			valueField.Set(reflect.ValueOf(int(value)))
		case reflect.String:
			var len uint32
			binary.Read(r, binary.LittleEndian, &len)

			dataRaw := make([]byte, len)
			binary.Read(r, binary.BigEndian, &dataRaw)

			valueField.SetString(string(dataRaw))
		default:
			return fmt.Errorf("bad type")
		}
	}
}