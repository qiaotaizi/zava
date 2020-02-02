package heap

func (c *Class) isAssignableFrom(class *Class) bool {
	s,t:=class,c
	if s==t{//同类型
		return true
	}

	if !t.IsInterface(){
		return s.isSubClassOf(t)//子类
	}else{
		return s.isImplements(t)//实现类
	}
}

func (cl *Class) isSubClassOf(other *Class) bool {
	for c:=cl.superClass;c!=nil;c=c.superClass{
		if c==other{
			return true
		}
	}

	return false
}

func (cl *Class)isImplements(iface *Class)bool{
	for c:=cl;c!=nil;c=c.superClass{
		for _,i:=range c.interfaces{
			if i==iface || i.isSubInterfaceOf(iface){
				return true
			}
		}
	}

	return false
}

func (cl *Class)isSubInterfaceOf(iface *Class)bool{
	for _,superInterface:=range cl.interfaces  {
		if superInterface==iface || superInterface.isSubInterfaceOf(iface){
			return true
		}
	}
	return false
}
