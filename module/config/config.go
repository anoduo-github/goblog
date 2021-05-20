package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var Cfg *ini.File

func Init(confPath string) {
	var err error
	Cfg, err = ini.Load(confPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		os.Exit(1)
	}
}
