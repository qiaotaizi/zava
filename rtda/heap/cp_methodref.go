package heap

import "github.com/qiaotaizi/zava/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef (cp *ConstantPool,refInfo *classfile.ConstantMethodrefInfo)*MethodRef{
	ref:=&MethodRef{}
	ref.cp=cp
	ref.copyMemberInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (m *MethodRef)ResolvedMethod()*Method{
	if m.method==nil{
		m.resolveMethodRef()
	}
	return m.method
}

//解析符号引用
func (m *MethodRef) resolveMethodRef() {
	d:=m.cp.class
	c:=m.ResolvedClass()//类符号引用
	if c.IsInterface(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	method:=lookupMethod(c,m.name,m.descriptor)
	if method==nil{
		panic("java.long.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}

	m.method=method
}

func lookupMethod(c *Class, name string, descriptor string) *Method {
	method:=LookupMethodInClass(c,name,descriptor)//在类和超类中查找
	if method==nil{
		//接口也会有默认方法
		method= lookupMethodInInterfaces(c.interfaces,name,descriptor)
	}
	return method
}
