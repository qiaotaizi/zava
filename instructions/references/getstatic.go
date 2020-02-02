package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//取得某个静态变量的值，推入操作数栈顶
type GET_STATIC struct {
	base.Index16Instruction
}

func (i *GET_STATIC) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	fieldRef:=cp.GetConstant(i.Index).(*heap.FieldRef)
	field:=fieldRef.ResolvedField()
	class:=field.Class()

	if !field.IsStatic(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor:=field.Descriptor()
	slotId:=field.SlotId()
	slots:=class.StaticVars()
	stack:=frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': //整形
		stack.PushInt(slots.GetInt(slotId))
	case 'F': //float
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J': //long
		stack.PushLong(slots.GetLong(slotId))
	case 'D': //double
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[': //引用，数组
		stack.PushRef(slots.GetRef(slotId))
	}
}

