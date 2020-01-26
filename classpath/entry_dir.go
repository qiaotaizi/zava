package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//读取目录下的类文件
type DirEntry struct {
	absDir string
}

//存放类目录绝对路径
func newDirEntry(path string) *DirEntry{
	absDir,err:=filepath.Abs(path)
	if err!=nil{
		panic(err)
	}

	return &DirEntry{absDir}
}

//读取文件
func (self *DirEntry)readClass(className string)([]byte ,Entry,error){
	fileName:=filepath.Join(self.absDir,className)
	data,err:=ioutil.ReadFile(fileName)
	return data,self,err
}

func (self * DirEntry)String() string{
	return self.absDir
}
