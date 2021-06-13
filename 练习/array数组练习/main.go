// 数组的练习

package main

import "fmt"

func main() {

	fmt.Println("####数组求和####")
	// 数组求和
	// 求数组 [1, 3, 5, 7, 8]的和
	var arr = [...]int{1, 3, 5, 7, 8}
	var sum int
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)

	fmt.Println("####求指定和的下标####")
	// 求数组 [1, 3, 5, 7, 8]的和为8的下标
	// 定义两个for循环，外层的从第一个开始遍历
	// 内层的for循环从外层后面的那个开始找
	// 他们的和为8
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf(`
下标：  %d   下标：  %d
对应值：%d   对应值：%d
和：    %d     +     %d  =  %d
`, i, j, arr[i], arr[j], arr[i], arr[j], arr[i]+arr[j])
			}
		}
	}
}
