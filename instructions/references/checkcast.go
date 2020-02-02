package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//类型强转指令
//(ClassA)obj
type CHECK_CAST struct {
	base.Index16Instruction
}

func (i *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	ref:=stack.PopRef()
	stack.PushRef(ref)

	if ref==nil{//如果对象是null，什么也不做
		return
	}

	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(i.Index).(*heap.ClassRef)
	class:=classRef.ResolvedClass()
	if !ref.IsInstanceOf(class){//类型转换失败，抛出异常
		panic("java.lang.ClassCastException")
	}


}

