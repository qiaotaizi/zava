package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//A instantOf B
//第一个操作数是根据索引获取的类信息，即B
//第二个操作数从操作数栈顶弹出，即A类实例的引用
//将判断结果推入操作数栈
//1：true 2：false
type INSTANCE_OF struct {
	base.Index16Instruction//类符号引用在常量池中的索引，用它获取B类
}

func (i *INSTANCE_OF) Execute(frame *rtda.Frame) {

	stack:=frame.OperandStack()
	ref:=stack.PopRef()
	if ref==nil{//如果A类实例是null，那么一定是false
		stack.PushInt(0)
		return
	}

	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(i.Index).(*heap.ClassRef)
	class:=classRef.ResolvedClass()
	if ref.IsInstanceOf(class){
		stack.PushInt(1)
	}else{
		stack.PushInt(0)
	}
}

