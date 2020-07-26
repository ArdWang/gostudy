package main

import "fmt"

//自己实现的算法
//var lastOccurred = make([] int, 0xffff)

func lengthOfNonRepeatingSubStr(s string) int {
	// stores last occurred pos +1
	// o means not seen
	//for i := range lastOccurred {
	//	lastOccurred[i] = 0
	//}

	//还是使用map
	lastOccurred := make(map[rune]int)

	start :=0
	maxLength :=0
	for i, ch := range [] rune(s){
		if lastI,ok := lastOccurred[ch];
			ok && lastI >= start{
			start = lastI+1
		}
		if i - start + 1 > maxLength{
			maxLength = i -start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("黑化肥挥发发会飞灰发合肥发黑会飞"))

	//Fileds,Split,Join
	//strings.
}
