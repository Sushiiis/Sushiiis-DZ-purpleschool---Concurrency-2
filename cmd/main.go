package main

import (
	"fmt"
	"math/rand"
)
//решение:
func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	var numbers []int

	//первая горутина
	go func(){
		for i := 0; i < 10; i++ {
			ri := rand.Intn(101)
			numbers = append(numbers, ri)
		}
		for _, num := range numbers{
		ch1 <- num
		}
	}()
	
	//вторая горутина
	go func(){
		for i := 0; i < 10; i++ {
			sum := <- ch1
			res := sum * sum
			ch2 <- res
		}
	}()

	for i := 0; i < 10; i++ {
   	 	dz := <- ch2
    	fmt.Println(dz)
	}

}