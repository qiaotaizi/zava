package base

import (
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame,method *heap.Method){
	thread:=invokerFrame.Thread()
	newFrame:=thread.NewFrame(method)
	thread.PushFrame(newFrame)//根据调用方法新建帧并推入栈顶

	argSlotCount:=int(method.ArgSlotCount())//参数传递，实例方法需要额外传递一个this引用
	if argSlotCount>0{
		for i:=argSlotCount-1;i>=0;i--{
			slot:=invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i),slot)
		}
	}
}


