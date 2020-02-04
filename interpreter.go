package main

import (
	"fmt"
	"github.com/qiaotaizi/zava/instructions"
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//解释器
//logInst参数用于控制是否将指令信息打印到控制台
func interpreter(method *heap.Method,logInst bool){
	//创建一个Thread实例，为方法创建一个帧，并把它推入虚拟机栈顶，最后执行方法
	thread:=rtda.NewThread()
	frame:=thread.NewFrame(method)
	thread.PushFrame(frame)

	//开启虚拟机指令处理循环
	defer catchErr(thread)
	loop(thread,logInst)
}

//计算pc，解码指令，执行指令
func loop(thread *rtda.Thread, logInst bool) {
	//逐一执行字节码指令
	reader:=&base.BytecodeReader{}
	for{
		frame:=thread.CurrentFrame()
		pc:=frame.NextPC()
		thread.SetPC(pc)

		//解码为指令
		reader.Reset(frame.Method().Code(),pc)
		opCode:=reader.ReadUint8()
		inst:= instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst{
			logInstruction(frame,inst)
		}

		//执行指令
		inst.Execute(frame)
		if thread.IsStackEmpty(){
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method:=frame.Method()
	className:=method.Class().Name()
	methodName:=method.Name()
	pc:=frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n",className,methodName,pc,inst,inst)
}

//目前虚拟机还不能优雅终止程序运行
//只能通过捕获ret指令返回的异常来终止程序
func catchErr(thread *rtda.Thread) {
	if r:=recover();r!=nil{
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty(){
		frame := thread.PopFrame()
		method:=frame.Method()
		className:=method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(),className,method.Name(),method.Descriptor())
	}
}
