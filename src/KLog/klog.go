package KLog

import (
	"fmt"
	"log"
)

func CheckErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		fmt.Println("err: ", err)
		return false
	}

	return true
}
