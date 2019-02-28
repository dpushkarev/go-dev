package main

import (
	"bytes"
	"reflect"
	"encoding/binary"
	"fmt"
)

type User struct {
	ID       int
	Name     string
	LastName string `unpack:"-"`
	Flag     int
}

func main() {
	data := []byte{
		128, 45, 34, 0,
		9, 0, 0, 0,
		120, 46, 120, 120, 120, 120, 120, 120, 120,
		10, 0, 0, 0,
	}

	u := new(User)
	err := UnpackReflect(u, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\n%#v", u)
}

func UnpackReflect(u interface{}, data []byte) error {
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
	return nil
}