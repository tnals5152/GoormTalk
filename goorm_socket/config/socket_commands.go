package config

import "fmt"

var funcMap = map[string]interface{}{
	"TestFunc": TestFunc,
}

type ResultJson struct {
	Command string
	Value   interface{}
}

func ConverToFunc(funcName string, value interface{}) {
	if function, ok := funcMap[funcName]; ok {
		function.(func(interface{}))(value)

	}
}

func TestFunc(value interface{}) {
	fmt.Println(value)
}
