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
