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
