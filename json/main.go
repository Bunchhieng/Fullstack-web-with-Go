package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
// ",omitempty" this field will not be outputted if it is zero value
//  "-" will not be outputted
// ",string" Go converts this field to its corresponding JSON type
type Server struct {
	ServerName string `json:"-"`
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	jsonPath := "/Users/Bunchhieng/Documents/Bunchhieng/gowork/src/github.com/Bunchhieng/Web/json/main.json"
	file, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("error: %v", err)
	}
	data, err := ioutil.ReadAll(file)
	s := ServerSlice{}
	json.Unmarshal(data, &s)
	fmt.Println(s)
}
