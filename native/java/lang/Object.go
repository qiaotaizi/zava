package lang

import (
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"unsafe"
)

func init() {

	className:="java/lang/Object"

	native.Register(className,"getClass",
		"()Ljava/lang/Class;",getClass)

	native.Register(className,"hashCode",
		"()I",hashCode)
	native.Register(className,"clone",
		"()Ljava/lang/Object;",clone)
}

//Object#getClass本地方法实现
func getClass(frame *rtda.Frame){
	this:=frame.LocalVars().GetThis()
	class:=this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

func hashCode(frame *rtda.Frame){
	this:=frame.LocalVars().GetThis()
	hash:=int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func clone(frame *rtda.Frame){
	this:=frame.LocalVars().GetThis()
	cloneable:=this.Class().Loader().LoadClass("java/lang/Cloneable")

	if !this.Class().IsImplements(cloneable){
		panic("java/lang/CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}
