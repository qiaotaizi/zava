package main

import (
	"fmt"
	"github.com/qiaotaizi/zava/classfile"
	"github.com/qiaotaizi/zava/classpath"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
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

//1,2,3章startJVM
//func startJVM(cmd *Cmd) {
//	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
//	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
//
//	className := strings.Replace(cmd.class, ".", "/", -1) //-1表示替换所有
//
//	//找到类并打印字节码
//	//classData, _, err := cp.ReadClass(className)
//	//if err != nil {
//	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
//	//	return
//	//}
//	//fmt.Printf("class data:%v\n", classData)
//
//	cf := loadClass(className, cp)
//	fmt.Println(cmd.class)
//	printCLassInfo(cf)
//}

//4章startJVM
//func startJVM(cmd *Cmd){
//	frame:=rtda.newFrame(100,100)
//	println("testLocalVars")
//	testLocalVars(frame.LocalVars())
//	println()
//	println("testOperandStack")
//	testOperandStack(frame.OperandStack())
//}

func testOperandStack(operandStack *rtda.OperandStack) {
	operandStack.PushInt(100)
	operandStack.PushInt(-100)
	operandStack.PushLong(2997924580)
	operandStack.PushLong(-2997924580)
	operandStack.PushFloat(3.1415926)
	operandStack.PushDouble(2.71828182845)
	operandStack.PushRef(nil)
	println("popRef",operandStack.PopRef())
	println("popDouble",operandStack.PopDouble())
	println("popFloat",operandStack.PopFloat())
	println("popLong",operandStack.PopLong())
	println("popLong",operandStack.PopLong())
	println("popInt",operandStack.PopInt())
	println("popInt",operandStack.PopInt())
}

func testLocalVars(localVars rtda.LocalVars) {
	localVars.SetInt(0,100)
	localVars.SetInt(1,-100)
	localVars.SetLong(2,2997924580)
	localVars.SetLong(4,-2997924580)
	localVars.SetFloat(6,3.1415926)
	localVars.SetDouble(7,2.71828182845)
	localVars.SetRef(9,nil)
	println("getInt",localVars.GetInt(0))
	println("getInt",localVars.GetInt(1))
	println("getLong",localVars.GetLong(2))
	println("getLong",localVars.GetLong(4))
	println("getFloat",localVars.GetFloat(6))
	println("getDouble",localVars.GetDouble(7))
	println("getReg",localVars.GetRef(9))
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printCLassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlag())
	fmt.Printf("this class: %v\n", cf.ThisClass())
	fmt.Printf("super class: %v\n", cf.SuperClass())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("	%s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, info := range cf.Methods() {
		fmt.Printf("	%s\n", info.Name())
	}
}

//5，6章startJVM
func startJVM(cmd *Cmd){
	cp:=classpath.Parse(cmd.xJreOption,cmd.cpOption)
	classLoader:=heap.NewClassLoader(cp,cmd.verboseClassFlag)
	className:=strings.Replace(cmd.class,".","/",-1)
	mainClass:=classLoader.LoadClass(className)
	mainMethod:=mainClass.GetMainMethod()
	//cf:=loadClass(className,cp)
	//mainMethod:=getMainMethod(cf)
	if mainMethod!=nil{
		interpreter(mainMethod,cmd.verboseInstFlag)
	}else{
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}
}

//在类文件中查找main方法
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	methods:=cf.Methods()
	for _,method:=range methods{
		if method.Name()=="main" && method.Descriptor()=="([Ljava/lang/String;)V"{
			return method

		}
	}
	return nil
}
