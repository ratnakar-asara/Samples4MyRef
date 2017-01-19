package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main() {
	defer wg.Wait()
        wg.Add(20)

	for i:=0;i<10;i++ {
	go func (){
		defer wg.Done();
		fmt.Printf(" Wrong way : %d \n", i)
	}()
	}
	
	for j:=0;j<10;j++ {
	go func (j int){
		defer wg.Done();
		fmt.Printf("Right Way : %d \n", j)
	}(j)
	}

}

