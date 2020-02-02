package heap

import "github.com/qiaotaizi/zava/classfile"

//字段引用

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool,refInfo *classfile.ConstantFieldrefInfo)*FieldRef{
	ref:=&FieldRef{}
	ref.cp=cp
	ref.copyMemberInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (f *FieldRef)ResolvedField()*Field{
	if f.field==nil{
		f.resolveFieldRef()
	}
	return f.field
}

func (f *FieldRef) resolveFieldRef() {
	d:=f.cp.class//当前类
	c:=f.ResolvedClass()//引用字段所在类
	field:=lookupField(c,f.name,f.descriptor)//在所在类中查找字段
	if field==nil{
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d){//引用字段对当前类可见
		panic("java.lang.IllegalAccessError")
	}
	f.field=field
}

func lookupField(c *Class, name string, descriptor string) *Field {
	//在当前类中查找字段
	for _,field:=range c.fields{
		//字段名相同且描述符相同
		if field.name==name && field.descriptor==descriptor{
			return field
		}
	}

	//在所有接口中查找字段，静态字段？
	for _,iface:=range c.interfaces{
		if field:=lookupField(iface,name,descriptor);field!=nil{
			return field
		}
	}

	if c.superClass!=nil{
		return lookupField(c.superClass,name,descriptor)
	}
	return nil
}

