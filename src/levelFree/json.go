package main

import(
	"fmt"
	"encoding/json"
)

type User struct{
	User_id int `json:"id,string"`
	Name string `json:"-"`
	Phone string
}

var jsonStr = `{"id": 13, "name": "Den", "phone": "8926-31-31"}`

func main()  {
	data := []byte(jsonStr)
	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("struct:\n\t%#v\n\n", u)

	u.Phone = `22132213`
	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Printf("json string:\n\t%#v\n\n", string(result))
}