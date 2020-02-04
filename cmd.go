package main

import (
	"flag"
	"fmt"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	xJreOption  string
	class       string
	args        []string
	verboseClassFlag bool //是否把类加载信息打印到控制台
	verboseInstFlag bool //是否把指令信息打印到控制台
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xJreOption, "Xjre", "", "path to jre")
	flag.BoolVar(&cmd.verboseClassFlag,"verbose:class",false,"print class load info to console")
	flag.BoolVar(&cmd.verboseInstFlag,"verbose:inst",false,"print instruction info to console")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Println("Usage: Zava [-options] class [args...]")
}
