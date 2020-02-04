package heap

//在类和超类中查找方法
func LookupMethodInClass(class *Class, name string, descriptor string) *Method {
	for c := class; c != nil; c = class.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

//在接口中查找方法
func lookupMethodInInterfaces(ifaces []*Class, name string, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}

		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}
