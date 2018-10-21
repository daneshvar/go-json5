package main

import (
	"fmt"
	"os"

	"github.com/daneshvar/json5"
)

func main() {
	if f, e := os.Open("./example.json5"); e == nil {
		defer f.Close()
		dec := json5.NewDecoder(f)

		var s interface{} // or struct
		if e := dec.Decode(&s); e == nil {
			fmt.Println(s)
		} else {
			fmt.Println(e)
		}
	} else {
		fmt.Println(e)
	}
}
