package main

import (
	"encry/config"
	"strings"
)

func Run() {

}

func loadOS(args []string, argLen int) {
	for i := 1; i < argLen; i++ {
		if ok := strings.Contains(args[i], "-"); ok {
			// other args
			switch args[i] {
			case "-l":
				i++
				config.LOCALPORT = args[i]
			case "-local":
				i++
				config.LOCALPORT = args[i]
			case "-r":
				i++
				config.REMOTEADDRESS = args[i]
			case "-remote":
				i++
				config.REMOTEADDRESS = args[i]
			case "-s": // 静默
				config.LOGSILENT = 1
			case "-h":
				help()
			}
		}
	}

}

func loadYaml() {

}
