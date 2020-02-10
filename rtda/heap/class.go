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
	initStarted bool//<clinit>是否已经开始执行
	jClass *Object//类的java.lang.Class对象
	sourceFile string//源代码文件名
}

func (c *Class)InitStarted()bool{
	return c.initStarted
}

func (c *Class)StartInit(){
	c.initStarted=true
}

func (c *Class)JClass()*Object{
	return c.jClass
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
	class.sourceFile=getSourceFile(cf)
	return class
}

//从属性表获取源代码信息
func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr:=cf.SourceFileAttributes();sfAttr!=nil{
		return sfAttr.FileName()
	}
	return "Unknown"
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

func (c *Class) Name() string {
	return c.name
}

func (c *Class) GetClinitMethod() *Method {
	return c.getStaticMethod("<clinit>","()V")
}

func (c *Class) Loader() *ClassLoader {
	return c.loader
}

func (c *Class) ArrayClass() *Class {
	arrayClassName:=getArrayClassName(c.name)
	return c.loader.loadArrayClass(arrayClassName)
}

func (c *Class) isJlObject() bool {
	return c.name=="java/lang/Object"
}

func (c *Class) isJlCloneable() bool {
	return c.name=="java/lang/Cloneable"
}

func (c *Class) isJioSerializable() bool {
	return c.name=="java/io/Serializable"
}

//根据字段名和描述符查找字段
func (c *Class) GetField(name string, descriptor string, isStatic bool) *Field {
	for class:=c;class!=nil;class=c.superClass{
		for _,field:=range c.fields{
			if field.IsStatic()==isStatic &&
				field.name==name &&
				field.descriptor==descriptor{
				return field
			}
		}
	}

	return nil
}

func (c *Class) JavaName() string {
	return strings.Replace(c.name,"/",".",-1)
}

func (c *Class) IsPrimitive() bool {
	_,ok:=primitiveTypes[c.name]
	return ok
}

func (c *Class) GetRefVar(name string, descriptor string) *Object {
	field:=c.GetField(name,descriptor,true)
	return c.staticVars.GetRef(field.slotId)
}

func (c *Class) SetRefVar(name,descriptor string,ref *Object){
	field:=c.GetField(name,descriptor,true)
	c.staticVars.SetRef(field.slotId,ref)
}

func (c *Class) GetInstanceMethod(name string, descriptor string) *Method {
	return c.getMethod(name,descriptor,false)
}

func (c *Class) getMethod(name string, descriptor string, isStatic bool) *Method {
	for class:=c;class!=nil;class=c.superClass{
		for _,method:=range class.methods{
			if method.IsStatic()==isStatic &&
				method.name==name &&
				method.descriptor==descriptor{
				return method
			}
		}
	}

	return nil
}

func (c *Class) SourceFile() string {
	return c.sourceFile
}
