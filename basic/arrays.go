package main

import "fmt"

//会 拷贝 数组
func printArray(arr [] int)  {
	arr[0] = 100
	for _,v := range arr{
		fmt.Println(v)
	}
}


func main() {

	var arr1 [5] int
	//arr2 :=[3]int{1,3,5}

	arr3 := [...]int{2,4,6,8,10}

	//var grid [4][5]int
	//
	//fmt.Println(arr1,arr2,arr3)
	//fmt.Println(grid)
	//
	//for i := 0; i<len(arr3); i++{
	//	fmt.Println(arr3[i])
	//}
	//
	//for _,v := range arr3{
	//	fmt.Println(v)
	//}

	printArray(arr1[:])
	printArray(arr3[:])
	//printArray(arr2)

	fmt.Println(arr1,arr3)

}
