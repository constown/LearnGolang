package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0、参数的校验
	// 0.1、传进来的 data 参数必须是指针类型（因为在函数中需要对其赋值）
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") // 新创建一个错误，返回一个err类型
		return
	}
	// 0.2、传进来的 data 参数必须是结构体类型指针，因为配置文件中各种键值对需要赋值给结构体中的字段
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct pointer") // 新创建一个错误，返回一个err类型
		return
	}
	// 1、读文件得到字节类型数据   "./conf.ini"
	b, err1 := ioutil.ReadFile(fileName)
	if err1 != nil {
		return err1
	}
	// 2、逐行读取数据
	lineSlice := strings.Split(string(b), "\n") // string(b)将字节类型转换成字符串
	var structName string
	for idx, line := range lineSlice {
		// 去掉空格
		line = strings.TrimSpace(line)
		// 如果是空行，就直接跳过
		if len(line) == 0 {
			continue
		}
		// 2.1、如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			// 2.2、如果是中括号（[）表示是节（section）
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx+1) // 新创建一个错误，返回一个err类型
				return
			}
			// 去掉中括号，去到中间去掉空格的内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1) // 新创建一个错误，返回一个err类型
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 找到了对应的结构体，把字段名记下来
					structName = field.Name
					//fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3、否则就是 = 分割的键值对
			// 用 = 分割，左边 是 key，右边 是 value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1) // 新创建一个错误，返回一个err类型
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//	根据 structName 去打他里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			structObjValue := v.Elem().FieldByName(structName) // 嵌套结构体的值信息
			structObjType := structObjValue.Type()             // 嵌套结构体的类型信息
			if structObjType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是个结构体", structName) // 新创建一个错误，返回一个err类型
				return
			}
			var filedName string
			var fileType reflect.StructField
			//	遍历嵌套结构体的每一个字段，判断 tag 是否等于这个 key， 等于就赋值
			for i := 0; i < structObjValue.NumField(); i++ {
				filed := structObjType.Field(i) // tag信息是存储在类型信息中的
				fileType = filed
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					filedName = filed.Name
					break
				}
			}
			if len(filedName) == 0 {
				// 在结构体中找不到对应的字段，就直接跳过
				continue
			}
			// 根据 filedName 去取出字段，赋值
			fileObj := structObjValue.FieldByName(filedName)
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line %d type error", idx+1) // 新创建一个错误，返回一个err类型
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var parseBool bool
				parseBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line %d type error", idx+1) // 新创建一个错误，返回一个err类型
					return
				}
				fileObj.SetBool(parseBool)
			case reflect.Float32, reflect.Float64:
				var parseFloat float64
				parseFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line %d type error", idx+1) // 新创建一个错误，返回一个err类型
					return
				}
				fileObj.SetFloat(parseFloat)
			}
		}

	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./exercise/iniParse/conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
	}
	fmt.Printf("%#v\n", cfg)
}
