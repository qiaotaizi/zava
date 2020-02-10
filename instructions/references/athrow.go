package references

import (
	"github.com/qiaotaizi/zava/instructions/base"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (*ATHROW) Execute(frame *rtda.Frame) {
	ex:=frame.OperandStack().PopRef()//异常对象引用

	if ex==nil{
		panic("java.lang.NullPointerException")
	}

	thread:=frame.Thread()
	if !findAndGotoExceptionHandler(thread,ex){
		handleUncaughtException(thread,ex)
	}
}

//遍历虚拟机栈后没有找到处理异常的pc
//由虚拟机打印栈信息
func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	//清空虚拟机栈，没有指令，虚拟机将停止运行
	thread.ClearStack()
	//获取Throwable对象的message字符串
	jMsg:=ex.GetRefVar("detailMessage","Ljava/lang/String;")
	goMsg:=heap.GoString(jMsg)
	println(ex.Class().JavaName()+": "+goMsg)

	//异常对象的extra属性存放的是虚拟机栈信息，打印之，属性设置参见Throwable.go
	stes:=reflect.ValueOf(ex.Extra())
	for i:=0;i<stes.Len();i++{
		ste:=stes.Index(i).Interface().(interface{
			String() string
		})
		println("\tat "+ste.String())
	}
}

//检查是否可以在虚拟机栈中找到可以处理异常的pc
func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame :=thread.CurrentFrame()
		pc:=frame.NextPC()-1//抛出异常的指令的pc

		handlerPc:=frame.Method().FindExceptionHandler(ex.Class(),pc)
		if handlerPc>0{
			stack:=frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPc)
			return true
		}

		thread.PopFrame()//当前方法不能处理异常，检查上一方法是否可以处理
		if thread.IsStackEmpty(){
			break
		}

	}

	return false
}



