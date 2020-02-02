package heap

//符号引用类常量的公共部分
type SymRef struct {
	cp *ConstantPool
	className string
	class *Class
}

func (s *SymRef)ResolvedClass()*Class{
	if s.class==nil{
		s.resolveClassRef()
	}
	return s.class
}

func (s *SymRef) resolveClassRef() {
	d:=s.cp.class//使用当前类的加载器加载引用类
	c:=d.loader.LoadClass(s.className)
	if ! c.isAccessibleTo(d){//检查引用类是否对当前类可见
		panic("java.lang.IllegalAccessError")
	}
	s.class=c
}
