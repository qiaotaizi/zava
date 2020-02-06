package heap

func (c *Class) isAssignableFrom(class *Class) bool {
	s,t:=class,c
	if s==t{//同类型
		return true
	}

	if !s.IsArray(){
		if !s.IsInterface(){
			if !t.IsInterface(){
				//对于非数组对象，s是class，t是class
				return s.IsSubClassOf(t)//判断s是t的子类
			}else{
				//对于非数组对象，s是class t是interface
				return s.IsImplements(t)//判断s是t的实现类
			}
		}else{
			if !t.IsInterface(){
				//对于非数组对象，s是interface，t是class
				return t.isJlObject()
			}else{
				return t.isSuperInterfaceOf(s)
			}
		}
	}else{
		if !t.IsArray(){
			if !t.IsInterface(){
				return t.isJlObject()
			}else{
				return t.isJlCloneable() || t.isJioSerializable()
			}
		}else{
			sc:=s.ComponentClass()
			tc:=s.ComponentClass()
			return sc==tc || tc.isAssignableFrom(sc)
		}
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

func (c *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(c)
}
