package rtda

type Frame struct {
	lower        *Frame        //链表前一节点
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
	thread *Thread//所属线程
	nextPC int//下个要执行的指令
}

func newFrame(thread *Thread,maxLocals,maxStack uint)*Frame{
	return &Frame{
		thread:       thread,
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
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
