package main

import (
	"acc/service"
	"fmt"
)

func main(){

	// file reader service for reading the values from file
	foodmap,err:=service.FileReader()

	if err != nil{
		fmt.Println("error while reading",err.Error())
		return
	}

	// findMax service used for finding the top 3 values
	max1,max2,max3:=service.Findmax(foodmap)

	fmt.Println("1. ",max1.Food)
	fmt.Println("2. ",max2.Food)
	fmt.Println("3. ",max3.Food)
}

