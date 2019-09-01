package main

import "fmt"

var version="version 0.0.1"

//install之后在goroot中查找二进制文件

func main() {
	cmd := parseCmd()
	if cmd.versionFlag{
		fmt.Println(version)
	}else if cmd.helpFlag || cmd.class==""{
		printUsage()
	}else{
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd){
	fmt.Printf("classpath: %s class: %s args: %v\n",cmd.cpOption,cmd.class,cmd.args)
}
