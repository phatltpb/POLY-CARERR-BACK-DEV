package myLog

import "fmt"

func Log(v interface{}) {
	fmt.Printf("%+v\n", v)

}
