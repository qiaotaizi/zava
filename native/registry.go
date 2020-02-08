package native

import "github.com/qiaotaizi/zava/rtda"

//本地方法注册与查找


//本地方法抽象，frame参数即是本地方法工作空间
type NativeMethod func(frame *rtda.Frame)

//本地方法注册哈希表
var registry=map[string]NativeMethod{}

//注册
func Register(className, methodName,methodDescriptor string,method NativeMethod){
	registry[concatNativeMethodKey(className,methodName,methodDescriptor)]=method
}

func concatNativeMethodKey(className,methodName,methodDescriptor string)string{
	return className+"~"+methodName+"~"+methodDescriptor
}

//查找
func FindNativeMethod(className,methodName,methodDescriptor string)NativeMethod{
	key:=concatNativeMethodKey(className,methodName,methodDescriptor)
	if method,ok:=registry[key];ok{
		return method
	}

	//java.lang.Object通过registerNatives方法来注册其他本地方法的
	if methodDescriptor=="()V" && methodName=="registerNatives"{
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(*rtda.Frame){
	//什么也不做
}