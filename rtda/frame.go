package rtda

type Frame struct {
	lower        *Frame        //链表前一节点
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		//局部变量表的大小和操作数栈的深度在编译期计算好
		//存储在class文件的method_info结构的Code属性中
	}
}

func (f *Frame) LocalVars()LocalVars{
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack{
	return f.operandStack
}
