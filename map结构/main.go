package main

import "fmt"

func main() {
	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	var m1 map[string]int         // 还没有初始化
	fmt.Println(m1 == nil)        // true
	m1 = make(map[string]int, 10) // 要估算好map容量，避免在程序运行期间再动态扩容
	m1["离线"] = 0
	m1["在线"] = 1
	fmt.Println("m1:", m1)
	fmt.Println("m1[\"离线\"]:", m1["离线"])

	value, ok := m1["隐身"] // 如果不存在这个key拿到对应值类型的零值
	if !ok {
		fmt.Println("无效的key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	for key := range m1 {
		fmt.Println(key)
	}

	// 删除
	delete(m1, "离线")
	delete(m1, "隐身")

	// map 和 slice 组合

	var s1 = make([]map[int]string, 0, 10)
	s1[0][100] = "A" // 越界
	fmt.Println(s1)
}
