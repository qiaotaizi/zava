# zava
自己动手写java虚拟机

第一章：  
实现zava查看命令选项

第二章：
实现在jre/lib、jre/ext、classpath目录下找类  
zava -Xjre=$JAVA_HOME/jre java.lang.Object  
或者在指定了JAVA_HOME环境变量之后  
zava java.lang.Object

第三章：  
解析class文件，实现一个简单版的javap命令  
zava -Xjre=$JAVA_HOME/jre java.lang.String  
或者  
zava java.lang.String

第四章：  
实现虚拟机栈  

第五章：  
实现大部分的虚拟机指令  
实现解释器  
现在可以在这个虚拟机上运行简单的java代码了！

第六章：  
实现类成员的引用  
现在虚拟机可以支持面向对象编程了  

第七章：  
实现方法调用  

第八章：  
实现数组初始化和java.lang.String初始化  

第九章：  
实现本地方法调用  

第十章：  
实现异常处理
