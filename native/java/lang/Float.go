package lang

import (
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"math"
)

func init() {
	className:="java/lang/Float"
	native.Register(className,"floatToRawIntBits",
		"(F)I",floatToRawIntBits)
	native.Register(className,"intBitsToFloat",
		"(I)F",intBitsToFloat)
}

func floatToRawIntBits(frame *rtda.Frame){
	value:=frame.LocalVars().GetFloat(0)
	bits:=math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

func intBitsToFloat(frame *rtda.Frame){
	bits:=frame.LocalVars().GetInt(0)
	value:=math.Float32frombits(uint32(bits))
	frame.OperandStack().PushFloat(value)
}
