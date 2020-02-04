package rtda

import "github.com/qiaotaizi/zava/rtda/heap"

type Frame struct {
	lower        *Frame        //链表前一节点
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
	thread *Thread//所属线程
	nextPC int//下个要执行的指令
	method *heap.Method
}

func newFrame(thread *Thread,method *heap.Method)*Frame{
	return &Frame{
		thread:       thread,
		localVars:newLocalVars(method.MaxLocals()),
		operandStack:newOperandStack(method.MaxStack()),
		method:method,
	}
}

func (f *Frame) LocalVars()LocalVars{
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack{
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) SetNextPC(pc int) {
	f.nextPC=pc
}

func (f *Frame) NextPC()int{
	return f.nextPC
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

//某些指令调用类的<clinit>方法时使用
//如果发现初始化方法尚未开始调用，开始调用初始化方法，指令执行到一半直接返回
//但程序计数器向下移动了
//这个方法使程序计数器回退一步
func (f *Frame) RevertNextPC() {
	f.nextPC=f.thread.pc
}
