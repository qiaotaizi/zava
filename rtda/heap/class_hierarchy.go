package heap

func (c *Class) isAssignableFrom(class *Class) bool {
	s,t:=class,c
	if s==t{//同类型
		return true
	}

	if !t.IsInterface(){
		return s.IsSubClassOf(t)//子类
	}else{
		return s.IsImplements(t) //实现类
	}
}

func (c *Class) IsSubClassOf(other *Class) bool {
	for class := c.superClass; class !=nil; class = class.superClass{
		if class ==other{
			return true
		}
	}

	return false
}

func (c *Class) IsSuperClassOf(other *Class)bool{
	return other.IsSubClassOf(c)
}

func (c *Class) IsImplements(iface *Class)bool{
	for class := c; class !=nil; class = class.superClass{
		for _,i:=range class.interfaces{
			if i==iface || i.isSubInterfaceOf(iface){
				return true
			}
		}
	}

	return false
}

func (c *Class)isSubInterfaceOf(iface *Class)bool{
	for _,superInterface:=range c.interfaces  {
		if superInterface==iface || superInterface.isSubInterfaceOf(iface){
			return true
		}
	}
	return false
}
