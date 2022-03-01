package config

import (
	"encoding/json"
	"fmt"

	"goorm_socket/utils"
)

var funcMap = map[string]interface{}{
	"TestFunc": TestFunc,
}

type ResultJson struct {
	Command string //functionName
	Value   interface{}
}

func ConverToFunc(jsonMessage []byte) {
	var resultJson ResultJson
	err := json.Unmarshal(jsonMessage, &resultJson)
	utils.ErrorCheck(err)

	if function, ok := funcMap[resultJson.Command]; ok {
		function.(func(interface{}))(resultJson.Value)
	}
}

func TestFunc(value interface{}) {
	fmt.Println(value)
}
