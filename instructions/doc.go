package instructions

//java虚拟机目前为字节码解释器提供了205条指令
//分为11类，分别是：
//常量
//加载
//存储
//操作数栈
//数学
//转换
//比较
//控制
//引用
//扩展
//保留（不允许出现在class文件中）

//这个包将实现其中9类指令的一部分

//解释器自身是一个循环
//当还有指令需要执行时进入一次循环
//在循环中：
//计算程序计数器
//解码指令
//执行指令
