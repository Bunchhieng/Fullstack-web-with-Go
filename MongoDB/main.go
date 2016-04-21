package main

import (
	"gopkg.in/mgo.v2" // go get gopkg.in/mgo.v2
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"syscall"
	"os"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("server1.example.com")
	checkErr(err)

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Database = test and collection = test
	c := session.DB("test").C("test")
	err = c.Insert(&Person{"Bunchhieng", "97843241"}, &Person{"Rany", "324923749"})
	checkErr(err)

	result := Person{}
	err = c.Find(bson.M{"name": "Bunchhieng"}).One(&result)
	checkErr(err)
	fmt.Println(result.Phone)
	if os.Interrupt == syscall.SIGINT {
		fmt.Println("User close the connection")
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
