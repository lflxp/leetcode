package main

import (
	"fmt"
)

func main() {

	// type Road struct {
	// 	Name   string
	// 	Number int
	// }
	// roads := []Road{
	// 	{"Diamond Fork", 29},
	// 	{"Sheep Creek", 51},
	// }

	// b, err := json.Marshal(roads)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var out bytes.Buffer
	// err = json.Indent(&out, b, "", "\t")

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// out.WriteTo(os.Stdout)

	// a := "abc (123)"
	// a := "abc 998"
	// fmt.Println(strings.Split(a, "(")[0])

	a := 123.123
	fmt.Println(fmt.Sprintf("%d", int(a)))
}
