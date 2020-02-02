package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//初始化类实例的指令

type NEW struct {
	base.Index16Instruction //类引用索引，从当前类的常量池拿到目标类信息，创建对象，推入操作数栈顶
}

func (n *NEW) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(n.Index).(*heap.ClassRef)
	class:=classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract(){
		panic("java.lang.InstantiationError")//接口和抽象类无法直接创建实例
	}

	ref :=class.NewObject()
	frame.OperandStack().PushRef(ref)
}









