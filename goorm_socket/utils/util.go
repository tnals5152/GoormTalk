package utils

import (
	"log"
	"runtime"
)

//에러 체크 후 에러 시 무조건 패닉 발생하는 함수 -> error발생 시 작동 X여서 다시 실행해야 될 때 사용
func IfErrorMakePanic(err error, errorInfo string) {
	if err != nil {
		log.Println(errorInfo)
		panic(err)
	}
}

func ErrorCheck(err error) {
	if err != nil {
		PrintErrFunc(err, 2)
	}
}

func PrintErrFunc(err error, depth int) {
	if err == nil {
		return
	}
	funcName, file, line, ok := runtime.Caller(depth)
	if ok {
		log.Printf("[%s] [%s] (%d line) - %s",
			file, runtime.FuncForPC(funcName).Name(), line, err.Error())
	}
}
