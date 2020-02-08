package lang

import (
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func init() {
	className:="java/lang/String"
	native.Register(className,"intern","()Ljava/lang/String;",intern)
}

func intern(frame *rtda.Frame)  {
	this:=frame.LocalVars().GetThis()
	interned:=heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
