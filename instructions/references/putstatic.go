package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//给静态变量赋值
//值从操作数栈中弹出
type PUT_STATIC struct {
	base.Index16Instruction //常量的符号引用
}

func (i *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()

	fieldRef := cp.GetConstant(i.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() { //final字段不允许被二次赋值,只能在初始化方法中进行赋值
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': //整形
		slots.SetInt(slotId, stack.PopInt())
	case 'F': //float
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J': //long
		slots.SetLong(slotId, stack.PopLong())
	case 'D': //double
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[': //引用，数组
		slots.SetRef(slotId, stack.PopRef())

	}

}
