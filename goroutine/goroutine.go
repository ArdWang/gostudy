package main

import (
	"fmt"
	"time"
)


func main() {
	//var a [10]int
	for i := 0; i<1000; i++ {
		go func(ii int) { //race
			for{
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
				//a[ii]++
				//runtime.Gosched()

			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}

