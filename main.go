package main

import (
	"encry/config"
	"encry/logs"
	"fmt"
	"log"
	"os"
)

func main() {

	welcome()

	args := os.Args
	argLen := len(os.Args)

	if argLen == 0 {
		loadYaml() // 通过yaml配置启动
	} else {
		loadOS(args, argLen) // 通过命令参数启动
	}

	// 验证
	if config.LOCALPORT == "" || config.REMOTEADDRESS == "" {
		log.Println("args is error")
		return
	}

	logs.Load()

	Run()

}

func help() {
	fmt.Println("+-----------------------------help information--------------------------------+")
	fmt.Println(`usage: "-listen port1 port2" #example: "gohtran -listen 8888 3389" `)
	fmt.Println(`       "-tran port1 ip:port2" #example: "gohtran -tran 8888 1.1.1.1:3389" `)
	fmt.Println(`       "-slave ip1:port1 ip2:port2" #example: "gohtran -slave 127.0.0.1:3389 1.1.1.1:8888" `)
	fmt.Println(`       "-e enable gzip and aes functionality`)
	fmt.Println(`       "-aes enable aes functionality, parameters is key, defaults to 16 bits`)
	fmt.Println(`       "-gzip enable gzip functionality`)
	fmt.Println(`       "-h -help program documentation`)
	fmt.Println(`       "-s -silent silent mode,no information is displayed`)
	fmt.Println(`       "-logs output transferred data to file`)
	fmt.Println(`============================================================`)
	fmt.Println("If you see xxxxxx, that means the data channel is established")
}

func welcome() {
	log.Println("============== welcome ==================")
	log.Println("Program execution begins...")
}
