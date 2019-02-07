package main

import (
	"fmt"
	"encoding/json"
)

var jsonAr = `[
		{"id": 13, "name": "Den", "phone": "8926-31-31"},
		{"last_name": "govnar", "address": "5th Avenue", "age": "3", "city": "Moscow"}
	]`

func main(){
	data := []byte(jsonAr)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("empty interface:\n\t%#v\n\n", user1)

}