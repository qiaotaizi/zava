package lang

import (
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func init(){
	className:="java/lang/System"
	native.Register(className,"arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V",arraycopy)
}

func arraycopy(frame *rtda.Frame){
	vars:=frame.LocalVars()
	src:=vars.GetRef(0)
	srcPos:=vars.GetInt(1)
	dest:=vars.GetRef(2)
	destPos:=vars.GetInt(3)
	length:=vars.GetInt(4)

	if src==nil || dest==nil{
		panic("java.lang.NullPointerException")
	}

	if !checkArrayCopy(src,dest){
		panic("java.lang.ArrayStoreException")
	}

	if srcPos<0 || destPos<0 || length<0 ||
		srcPos+length>src.ArrayLength() ||
		destPos+length>dest.ArrayLength(){
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src,dest,srcPos,destPos,length)

}

//执行拷贝前检查数组兼容
func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {
	srcClass:=src.Class()
	destClass:=dest.Class()
	if !srcClass.IsArray() || !destClass.IsArray(){
		return false
	}

	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive(){
		return srcClass==destClass
	}

	return true
}
