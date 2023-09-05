package models

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Config *ini.File

func init() {
	Config, err = ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}
