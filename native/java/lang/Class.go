package lang

import (
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func init() {

	className:="java/lang/Class"

	native.Register(className,"getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;",getPrimitiveClass)
	native.Register(className,"getName0",
		"()Ljava/lang/String;",getName0)
	native.Register(className,"desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z",desiredAssertionStatus0)
}


//获取对象关联的Class对象的名称
func getName0(frame *rtda.Frame){
	this:=frame.LocalVars().GetThis()
	class:=this.Extra().(*heap.Class)

	name:=class.JavaName()

	nameObj:=heap.JString(class.Loader(),name)
	frame.OperandStack().PushRef(nameObj)
}

//获取基本数据类型的Class对象
func getPrimitiveClass(frame *rtda.Frame){
	nameObj:=frame.LocalVars().GetRef(0)//java方法中的name参数
	name:=heap.GoString(nameObj)

	loader:=frame.Method().Class().Loader()
	class:=loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

func desiredAssertionStatus0(frame *rtda.Frame)  {
	//不详细讨论断言，这里简单地将false推入栈顶
	frame.OperandStack().PushBoolean(false)
}
