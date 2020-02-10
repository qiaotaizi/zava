package lang

import (
	"github.com/qiaotaizi/zava/native"
	"github.com/qiaotaizi/zava/rtda"
	"github.com/qiaotaizi/zava/rtda/heap"
)

func init() {
	className:="java/lang/Throwable"
	native.Register(className,"fillInStackTrace",
		"(I)Ljava/lang/Throwable;",fillInStackTrace)
}

//抛出异常时读取java虚拟机栈
func fillInStackTrace(frame *rtda.Frame){
	this:=frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes:=createStackElements(this,frame.Thread())
	this.SetExtra(stes)
}

func createStackElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	//栈顶两帧分别是fillInStackTrace(int)和fillInStackTrace()方法，跳过这两帧
	//这两帧下面的几个帧正在执行异常类的构造函数，需要distanceToObject来计算具体跳过几帧
	skip:=distanceToObject(tObj.Class())+2
	//获取虚拟机栈
	frames:=thread.GetFrames()[skip:]
	stes:=make([]*StackTraceElement,len(frames))
	for i,frame:=range frames{
		stes[i]=createStackElement(frame)
	}
	return stes
}

//根据帧创建StackTraceElement
func createStackElement(frame *rtda.Frame) *StackTraceElement {
	method:=frame.Method()
	class:=method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC()-1),
	}
}

func distanceToObject(class *heap.Class) int {
	distance:=0
	for c:=class.SuperClass();c!=nil;c=c.SuperClass(){
		distance++
	}
	return distance
}

type StackTraceElement struct {
	fileName string
	className string
	methodName string
	lineNumber int  //帧正在执行的代码行数
}
