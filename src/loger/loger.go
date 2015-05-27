package loger

import (
	. "fmt"
)

var D *bool

func Info(info ...interface{}) {
	Println(info)
}

func ErrMsg(info ...interface{}) {
	Println(info)
}

func Debug(debug ...interface{}) {
	if *D {
		Println(debug)
	}
}
