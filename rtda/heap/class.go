package heap

import (
	"github.com/qiaotaizi/zava/classfile"
	"strings"
)

type Class struct {
	accessFlags uint16
	name string//类名，全限定名，形如java/lang/Object
	superClassName string
	interfaceNames []string
	constantPool *ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader //类加载器指针
	superClass *Class //超类指针
	interfaces []*Class //接口指针
	instanceSlotCount uint //实例变量占据的空间大小
	staticSlotCount uint //类变量占据的空间大小
	staticVars Slots //静态变量(多线程共享变量)
}

//将类文件转换为类对象
func newClass(cf *classfile.ClassFile)*Class{
	class :=&Class{}

	class.accessFlags=cf.AccessFlag()
	class.name=cf.ClassName()
	class.superClassName=cf.SuperClassName()
	class.interfaceNames=cf.InterfaceNames()
	class.constantPool=newConstantPool(class,cf.ConstantPool())
	class.fields=newFields(class,cf.Fields())
	class.methods=newMethods(class,cf.Methods())
	return class
}

func (c *Class)IsPublic() bool{
	return 0!=c.accessFlags & ACC_PUBLIC
}
func (c *Class)IsFinal()bool{
	return 0!=c.accessFlags&ACC_FINAL
}
func (c *Class)IsSuper()bool{
	return 0!=c.accessFlags&ACC_SUPER
}
func (c *Class)IsInterface()bool{
	return 0!=c.accessFlags&ACC_INTERFACE
}
func (c *Class)IsAbstract()bool{
	return 0!=c.accessFlags&ACC_ABSTRACT
}
func (c *Class)IsSynthetic()bool{
	return 0!=c.accessFlags&ACC_SYNTHETIC
}
func (c *Class)IsAnnotation()bool{
	return 0!=c.accessFlags&ACC_ANNOTATION
}
func (c *Class)IsEnum()bool{
	return 0!=c.accessFlags&ACC_ENUM
}

func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() || c.GetPackageName()==other.GetPackageName()
}

func (c *Class) GetPackageName() string {
	if i:=strings.LastIndex(c.name,"/");i>=0{
		return c.name[:i]
	}
	return ""
}

func (c *Class) NewObject() *Object {
	return newObject(c)
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main","([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name string, descriptor string) *Method {
	for _,method:=range c.methods{
		if method.IsStatic() &&
			method.name==name && method.descriptor==descriptor{
			return method
		}
	}
	return nil
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) Name() interface{} {
	return c.name
}

func newObject(c *Class) *Object {
	return &Object{
		fields:newSlots(c.instanceSlotCount),
		class:c,
	}
}
