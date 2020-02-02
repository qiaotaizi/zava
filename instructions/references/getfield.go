package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

type GET_FIELD struct {
	base.Index16Instruction
}

func (i *GET_FIELD) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	fieldRef:=cp.GetConstant(i.Index).(*heap.FieldRef)
	field:=fieldRef.ResolvedField()
	if field.IsStatic(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack:=frame.OperandStack()
	ref:=stack.PopRef()
	if ref==nil{
		panic("java.lang.NullPointerException")
	}

	descriptor:=field.Descriptor()
	slotId:=field.SlotId()
	slots:=ref.Fields()

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

