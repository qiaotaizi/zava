package reserved

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/native"
	_ "github.com/qiaotaizi/zava/native/java/lang" //调用包中的init函数
	_ "github.com/qiaotaizi/zava/native/sun/misc"
	"github.com/qiaotaizi/zava/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (*INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame) //本地方法调用
}
