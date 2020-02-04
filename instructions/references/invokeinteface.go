package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

//调用接口方法
//接口方法理论上也可以使用invokevirtual进行调用
//但是由于java是单继承多接口实现，invokevirtaul使用了虚方法表加速超类实例方法的查找
//接口方法则不能使用这个技术提高查找效率，所以做了指令层面的区分
type INVOKE_INTERFACE struct {
	index uint
}

func (i *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	i.index=uint(reader.ReadUInt16())//方法引用的常量池索引
	reader.ReadUint8()//slotCount,但是由于可以直接计算，这里忽略这个值
	reader.ReadUint8()//用于后向兼容的字节，总是0，这里忽略它
}

func (i INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	methodRef:=cp.GetConstant(i.index).(*heap.InterfaceMethodRef)
	resolvedMethod:=methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate(){
		//java8之后实际上接口可以定义static方法和private（java9）方法，但是这里不讨论这种情况
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref:=frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)
	if ref==nil{
		panic("java.lang.NullPointerException")
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()){
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked:=heap.LookupMethodInClass(ref.Class(),methodRef.Name(),methodRef.Descriptor())
	if methodToBeInvoked==nil || methodToBeInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame,methodToBeInvoked)

}

