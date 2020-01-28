package main

import (
	"fmt"
	"github.com/qiaotaizi/zava/classfile"
	"github.com/qiaotaizi/zava/classpath"
	"strings"
)

var version = "version 0.0.1"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println(version)
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1) //-1表示替换所有

	//找到类并打印字节码
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	//	return
	//}
	//fmt.Printf("class data:%v\n", classData)

	cf:=loadClass(className,cp)
	fmt.Println(cmd.class)
	printCLassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile{
	classData,_,err:=cp.ReadClass(className)
	if err!=nil{
		panic(err)
	}
	cf,err:=classfile.Parse(classData)
	if err!=nil{
		panic(err)
	}
	return cf
}

func printCLassInfo(cf *classfile.ClassFile){
	fmt.Printf("version: %v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count: %v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n",cf.AccessFlag())
	fmt.Printf("this class: %v\n",cf.ThisClass())
	fmt.Printf("super class: %v\n",cf.SuperClass())
	fmt.Printf("interfaces: %v\n",cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _,f := range cf.Fields() {
		fmt.Printf("	%s\n",f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, info := range cf.Methods() {
		fmt.Printf("	%s\n",info.Name())
	}
}
