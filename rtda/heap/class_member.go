package heap

import "github.com/qiaotaizi/zava/classfile"

//字段和方法的共有部分
type ClassMember struct {
	accessFlags uint16
	name string
	descriptor string
	class *Class
}

//从类文件中拷贝成员信息
func (cm *ClassMember)copyMemberInfo(memberInfo *classfile.MemberInfo){
	cm.accessFlags=memberInfo.AccessFlag()
	cm.name=memberInfo.Name()
	cm.descriptor=memberInfo.Descriptor()
}

func (cm *ClassMember)IsPublic()bool{
	return 0!=cm.accessFlags & ACC_PUBLIC
}
func (cm *ClassMember)IsPrivate()bool{
	return 0!=cm.accessFlags & ACC_PRIVATE
}
func (cm *ClassMember)IsProtected()bool{
	return 0!=cm.accessFlags & ACC_PROTECTED
}
func (cm *ClassMember)IsStatic()bool{
	return 0!=cm.accessFlags & ACC_STATIC
}
func (cm *ClassMember)IsFinal()bool{
	return 0!=cm.accessFlags & ACC_FINAL
}
func (cm *ClassMember)isSynthetic()bool{
	return 0!=cm.accessFlags & ACC_SYNTHETIC
}

func (cm *ClassMember)Descriptor()string{
	return cm.descriptor
}

func (cm *ClassMember) isAccessibleTo(other *Class) bool{
	//类成员是公开的，其他类可以引用
	if cm.IsPublic(){
		return true
	}
	c:=cm.class//成员所属类
	//protected
	if cm.IsProtected(){
		return other==c || other.isSubClassOf(c) || c.getPackageName()==other.getPackageName()
	}
	//friendly
	if !cm.IsPrivate(){
		return c.getPackageName()==other.getPackageName()
	}
	//private
	return other==c
}
