package heap

import "github.com/qiaotaizi/zava/classfile"

//接口方法引用

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo)*InterfaceMethodRef{
	ref :=&InterfaceMethodRef{}
	ref.cp=cp
	ref.copyMemberInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (i *InterfaceMethodRef)ResolvedInterfaceMethod() *Method{
	if i.method==nil{
		i.resolveInterfaceMethodRef()
	}
	return i.method
}

func (i *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d:=i.cp.class//当前类
	c:=i.ResolvedClass()//引用方法所属类

	if !c.IsInterface(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	method:=lookupInterfaceMethod(c,i.name,i.descriptor)

	if method==nil{
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}

	i.method=method


}

func lookupInterfaceMethod(iface *Class, name string, descriptor string) *Method {
	for _,method:=range iface.methods{
		if method.name==name && method.descriptor==descriptor{
			return method
		}
	}

	return lookupMethodInInterfaces(iface.interfaces,name,descriptor)
}


