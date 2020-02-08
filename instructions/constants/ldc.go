package constants

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//ldc从运行时常量池中加载常量值并推入操作数栈顶
//ldc=LoaD Constant?

type LDC struct {
	base.Index8Instruction
}

func (i *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame,i.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (i *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame,i.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack:=frame.OperandStack()
	class:=frame.Method().Class()
	c:=class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedString :=heap.JString(class.Loader(),c.(string))
		stack.PushRef(internedString)
	case *heap.ClassRef:
		classRef:=c.(*heap.ClassRef)
		classObj:=classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	//case MethodType,MethodHandle:
	default:
		panic("todo: ldc!")
	}
}

//处理long和double
type LDC2_W struct {
	base.Index16Instruction
}

func (i *LDC2_W) Execute(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	cp:=frame.Method().Class().ConstantPool()
	c:=cp.GetConstant(i.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

