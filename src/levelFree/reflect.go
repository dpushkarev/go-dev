package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int
	PastName string
	Flags    int
}

func main() {
	u := &User{
		ID:       42,
		PastName: "Den",
		Flags:    33,
	}
	err := PrintReflect(u)
	if err != nil {
		panic(err)
	}
}

func PrintReflect(u interface{}) error {
	val := reflect.ValueOf(u).Elem()
	fmt.Printf("%T have %d fields\n\n", u,val.NumField())
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		typeFuield := val.Type().Field(i)
		fmt.Printf("Name: %v, Value: %v, Type: %v, Tag: %v\n\n",
			typeFuield.Name,
			value,
			typeFuield.Type,
			typeFuield.Tag,
		)
	}
	return nil
}
