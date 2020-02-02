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
