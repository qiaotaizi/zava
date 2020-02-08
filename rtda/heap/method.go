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
	for i, cfMethod :=range cfMethods{
		methods[i]=newMethod(class,cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {

	method:=&Method{}
	method.class=class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md:=parseMethodDescriptor(method.descriptor)
	method.calArgSlotCount(md.parameterTypes)
	if method.IsNative(){
		method.injectAttributeCode(md.returnType) //注入本地方法字节码
	}
	return method
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
func (m *Method) calArgSlotCount(parameterTypes []string) {
	for _,paramType:=range parameterTypes{
		m.argSlotCount++
		if paramType=="J" || paramType=="D"{
			m.argSlotCount++//long或double占两个槽
		}
	}

	if !m.IsStatic(){
		m.argSlotCount++//给this引用一个位置
	}


}

//本地方法字节码注入
func (m *Method) injectAttributeCode(returnType string) {
	m.maxStack=4//TODO 暂时写死为4
	m.maxLocals=m.argSlotCount//本地方法帧的局部变量表用来存放参数值
	switch returnType[0] {
	case 'V':
		m.code=[]byte{0xfe,0xb1}//return
	case 'D':
		m.code=[]byte{0xfe,0xaf}//dreturn
	case 'F':
		m.code=[]byte{0xfe,0xae}//freturn
	case 'J':
		m.code=[]byte{0xfe,0xad}//lreturn
	case 'L','[':
		m.code=[]byte{0xfe,0xb0}//areturn
	default:
		m.code=[]byte{0xfe,0xac}//ireturn
	}
}
