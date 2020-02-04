package heap

import (
	"github.com/qiaotaizi/zava/classfile"
)

//类的方法
type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
	argSlotCount uint
}

func (m *Method) ArgSlotCount()uint{
	return m.argSlotCount
}

//maxStack maxLocals code等信息
//属于属性信息，从类文件中拷贝之
func (m *Method) copyAttributes(method *classfile.MemberInfo) {
	if codeAttr:=method.CodeAttribute();codeAttr!=nil{
		m.maxStack=codeAttr.MaxStack()
		m.maxLocals=codeAttr.MaxLocals()
		m.code=codeAttr.Code()
	}
}

//初始化方法表
func newMethods(class *Class ,cfMethods []*classfile.MemberInfo)[]*Method{
	methods:=make([]*Method,len(cfMethods))
	for i,method:=range cfMethods{
		methods[i]=&Method{}
		methods[i].class=class
		methods[i].copyMemberInfo(method)
		methods[i].copyAttributes(method)
		methods[i].calArgSlotCount()
	}
	return methods
}

func (m *Method) IsSynchronized()bool{
	return 0!=m.accessFlags & ACC_SYNCHRONIZED
}

func (m *Method)IsBridge()bool{
	return 0!=m.accessFlags & ACC_BRIDGE
}

func (m *Method)IsVarargs()bool{
	return 0!=m.accessFlags & ACC_VARARGS
}

func (m *Method) IsNative()bool{
	return 0!=m.accessFlags & ACC_NATIVE
}

func (m *Method) IsAbstract() bool{
	return 0!=m.accessFlags & ACC_ABSTRACT
}

func (m *Method) IsStrict() bool{
	return 0!=m.accessFlags & ACC_STRICT
}

func (m *Method) Class() *Class {
	return m.class
}

func (m *Method) Name() string {
	return m.name
}

func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

func (m *Method)Code()[]byte{
	return m.code
}

//计算方法参数数量
func (m *Method) calArgSlotCount() {
	parsedDescriptor:=parseMethodDescriptor(m.descriptor)
	for _,paramType:=range parsedDescriptor.parameterTypes{
		m.argSlotCount++
		if paramType=="J" || paramType=="D"{
			m.argSlotCount++//long或double占两个槽
		}
	}

	if !m.IsStatic(){
		m.argSlotCount++//给this引用一个位置
	}


}
