package heap

import (
	"fmt"
	"github.com/qiaotaizi/zava/classfile"
	"github.com/qiaotaizi/zava/classpath"
)

//类加载器
type ClassLoader struct {
	cp *classpath.Classpath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath)*ClassLoader{
	return &ClassLoader{
		cp:cp,
		classMap:make(map[string]*Class),
	}
}

//类加载器根据类名加载类数据至方法区，并返回Class对象
func (l *ClassLoader) LoadClass(name string) *Class{
	if class,ok:=l.classMap[name];ok{
		return class//类已经加载
	}
	return l.loadNonArrayClass(name)
}

//加载非数组类
func (l *ClassLoader) loadNonArrayClass(name string) *Class {
	data,entry:=l.readClass(name)//读取文件数据
	class:=l.defineClass(data)//解析生成类，放入方法区
	link(class)//链接
	fmt.Printf("[Loaded %s from %s]\n",name,entry)
	return class
}

//链接分为两个阶段：
//验证和准备
func link(class *Class) {
	verify(class)
	prepare(class)
}

func prepare(class *Class) {
	calcInstanceFieldSlotsIds(class) //查询实例变量数量，并给它们编号
	calcStaticFieldSlotsIds(class)   //查询静态变量数量，并给它们编号
	allocAndInitStaticVars(class)    //为静态变量分配空间并执行初始化
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars=newSlots(class.staticSlotCount)
	for _,field:=range class.fields{
		if field.IsStatic() && field.IsFinal(){
			initStaticFinalVar(class,field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars:=class.staticVars
	cp:=class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId:=field.slotId
	if cpIndex>0{
		switch field.Descriptor() {
		case "Z","B","C","S","I"://布尔，字符，整数等
			val:=cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId,val)
		case "J"://long
			val:=cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId,val)
		case "F"://float
			val:=cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId,val)
		case "D"://double
			val:=cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId,val)
		case "Ljava/lang/String;"://字符串
			panic("todo")
		}
	}
}

func calcStaticFieldSlotsIds(class *Class) {
	slotId:=uint(0)
	for _,field:=range class.fields{
		if field.IsStatic(){
			field.slotId=slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}
	class.staticSlotCount=slotId
}

func calcInstanceFieldSlotsIds(class *Class) {
	slotId:=uint(0)
	if class.superClass!=nil{
		slotId=class.superClass.instanceSlotCount
	}
	for _,field :=range class.fields{
		if !field.IsStatic(){
			field.slotId=slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++//long或double占两个位置
			}
		}
	}
	class.instanceSlotCount=slotId
}

//验证暂不实现
func verify(class *Class) {
	// TODO
}

func (l *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data,entry,err:=l.cp.ReadClass(name)
	if err!=nil{
		panic("java.lang.ClassNotFoundException: "+name)
	}
	return data,entry
}

func (l *ClassLoader) defineClass(data []byte) *Class {
	class:=parseClass(data)
	class.loader=l
	resolveSuperClass(class)
	resolveInterfaces(class)
	l.classMap[class.name]=class
	return class
}

func resolveInterfaces(class *Class) {
	interfaceCount:= len(class.interfaceNames)
	if interfaceCount>0{
		class.interfaces=make([]*Class,interfaceCount)
		for i,interfaceName:=range class.interfaceNames{
			class.interfaces[i]=class.loader.LoadClass(interfaceName)
		}
	}
}

func resolveSuperClass(class *Class) {
	if class.name!= "java/lang/Object"{
		class.superClass=class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf,err:=classfile.Parse(data)
	if err!=nil{
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

