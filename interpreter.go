package main

import (
	"fmt"
	"github.com/qiaotaizi/zava/classfile"
	"github.com/qiaotaizi/zava/instructions"
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
)

//解释器

func interpreter(methodInfo *classfile.MemberInfo){
	//获取方法字节码信息
	codeAttr:=methodInfo.CodeAttribute()
	maxLocals:=codeAttr.MaxLocals()
	maxStack:=codeAttr.MaxStack()
	bytecode:=codeAttr.Code()
	//创建一个Thread实例，为方法创建一个帧，并把它推入虚拟机栈顶，最后执行方法
	thread:=rtda.NewThread()
	frame:=thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)

	//开启虚拟机指令处理循环
	defer catchErr(frame)
	loop(thread,bytecode)
}

//计算pc，解码指令，执行指令
func loop(thread *rtda.Thread, bytecode []byte) {
	//逐一执行字节码指令
	frame:=thread.PopFrame()
	reader:=&base.BytecodeReader{}
	for{
		pc:=frame.NextPC()
		thread.SetPC(pc)

		//解码为指令
		reader.Reset(bytecode,pc)
		opCode:=reader.ReadUint8()
		inst:= instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//执行指令
		fmt.Printf("pc:%2d inst:%T %v\n",pc,inst,inst)
		inst.Execute(frame)
	}
}

//目前虚拟机还不能优雅终止程序运行
//只能通过捕获ret指令返回的异常来终止程序
func catchErr(frame *rtda.Frame) {
	if r:=recover();r!=nil{
		fmt.Printf("Localvars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}
