package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//给实例变量赋值
//需要三个操作数：
//常量池索引-指令自带
//变量值-操作数栈顶
//对象引用-操作数栈顶
type PUT_FIELD struct {
	base.Index16Instruction
}

func (i *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod:=frame.Method()
	currentClass:=currentMethod.Class()
	cp:=currentClass.ConstantPool()
	fieldRef:=cp.GetConstant(i.Index).(*heap.FieldRef)
	field:=fieldRef.ResolvedField()

	if field.IsStatic(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal(){
		if currentClass!=field.Class() || currentMethod.Name()!="<init>"{
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor:=field.Descriptor()
	slotId:=field.SlotId()
	stack:=frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': //整形
		val:=stack.PopInt()//变量值
		ref:=stack.PopRef()//对象引用
		if ref==nil{
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId,val)
	case 'F': //float
		val:=stack.PopFloat()
		ref:=stack.PopRef()
		if ref==nil{
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId,val)
	case 'J': //long
		val:=stack.PopLong()
		ref:=stack.PopRef()
		if ref==nil{
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId,val)
	case 'D': //double
		val:=stack.PopDouble()
		ref:=stack.PopRef()
		if ref==nil{
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId,val)
	case 'L', '[': //引用，数组
		val:=stack.PopRef()
		ref:=stack.PopRef()
		if ref==nil{
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId,val)
	}

}

