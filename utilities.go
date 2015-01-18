package main

import (
	"fmt"
	_ "net/http"
)

func returnErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
