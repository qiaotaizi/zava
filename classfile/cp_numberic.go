package classfile

import (
	"fmt"
	"math"
)

//4字节整型常量
type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) Value() int32 {
	return c.val
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	c.val = float32(bytes)
}

func (c *ConstantFloatInfo) Value() float32 {
	return c.val
}

type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	c.val = int64(bytes)
}

func (c *ConstantLongInfo) Value() int64 {
	return c.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	fmt.Printf("ConstantDoubleInfo#readInfo float64转型结果%f,math.Float64frombits转型结果%f", float64(bytes), math.Float64frombits(bytes))
	c.val = math.Float64frombits(bytes)
}

func (c *ConstantDoubleInfo) Value() float64 {
	return c.val
}
