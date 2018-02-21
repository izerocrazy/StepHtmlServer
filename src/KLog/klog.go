package KLog

import (
	"fmt"
	"log"
	"runtime/debug"
)

func CheckErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		fmt.Println("err: ", err)
		return false
	}

	return true
}

func Asset(value bool, info string, a ...interface{}) bool {
	if value == false {
		log.Println(info, a)
		debug.PrintStack()
		log.Fatal("Asset Fail")
		return false
	}

	return true
}

func Log(szFormat string, a ...interface{}) {
	log.Printf(szFormat, a...)
}
