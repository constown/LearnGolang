package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

// TypeOf 可以拿到类型相关的信息
func reflectType(x interface{}) {
	// 查看传进来的 x 是什么类型
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	// type name 和 type kind
	// 可以理解为 name 是自己造的 "表面类型"， kind 表示 "原始类型"
	// 比如 name 对应 Cat， kind 对应 struct
	fmt.Printf("type-name:%v type-kind:%v\n", v.Name(), v.Kind())
}

// ValueOf 可以拿到value相关的信息
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() // 值的类型种类
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值,反射传指针的时候一定要调用这个方法
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// isNil 常用于判断指针是否为空，IsValid 常用于判断返回值时候有效

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
	var c = Cat{}
	reflectType(c) // type-name:Cat type-kind:struct
	// ValueOf
	reflectValue(a)
	reflectSetValue2(&b)
	fmt.Println(b)
}
