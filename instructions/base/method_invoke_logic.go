package base

import (
	"fmt"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame,method *heap.Method){
	thread:=invokerFrame.Thread()
	fmt.Println("new frame of method:",method.Name(),"method arg slot count=",method.ArgSlotCount())
	newFrame:=thread.NewFrame(method)
	thread.PushFrame(newFrame)//根据调用方法新建帧并推入栈顶

	argSlotCount:=int(method.ArgSlotCount())
	if argSlotCount>0{
		for i:=argSlotCount-1;i>=0;i--{
			//待调用方法的参数从方法调用者所处的帧操作数栈中弹出，推入待调用方法帧的局部变量表中
			slot:=invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i),slot)
		}
	}
}


