package misc

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func init() {
	className:="sun/misc/VM"
	native.Register(className,"initialize","()V",initialize)
}

func initialize(frame *rtda.Frame){
	vmClass:=frame.Method().Class()
	savedProps:=vmClass.GetRefVar("savedProps","Ljava/util/Properties;")
	key:=heap.JString(vmClass.Loader(),"foo")
	val:=heap.JString(vmClass.Loader(),"bar")
	frame.OperandStack().PushRef(savedProps)

	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)

	propsClass:=vmClass.Loader().LoadClass("java/lang/Properties")
	setPropMethod:=propsClass.GetInstanceMethod("setProperty","(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame,setPropMethod)
}
