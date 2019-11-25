package info

import (
	"fmt"
	"reflect"
)

func Trace(value interface{}, varName string) {
	bar := "\n--- " + varName + " -------------------------------------------------------------"
	valueOf := fmt.Sprintf("Value of \"%s\": %v", varName, value)
	typeOf := fmt.Sprintf("Type of \"%s\": %v", varName, reflect.TypeOf(value))

	str := fmt.Sprintln(bar)
	str += fmt.Sprintln(valueOf)
	str += fmt.Sprintln(typeOf)
	fmt.Println(str)
}

func ConsoleMsg(s string) {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(s)
}
