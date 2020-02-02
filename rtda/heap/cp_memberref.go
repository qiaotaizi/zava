package heap

import "github.com/qiaotaizi/zava/classfile"

//类成员引用(字段，方法)

type MemberRef struct {
	SymRef
	name string
	descriptor string
}

func (m *MemberRef) copyMemberInfo(refInfo *classfile.ConstantMemberrefInfo){
	m.className=refInfo.ClassName()
	m.name,m.descriptor=refInfo.NameAndDescriptor()
}

func (m *MemberRef) Name()string{
	return m.name
}

func (m *MemberRef)Descriptor()string{
	return m.descriptor
}


