package service

import "acc/models"

func Findmax(food_map map[string]int) (models.MaxCounter,models.MaxCounter,models.MaxCounter){

	// logic for finding the top 3 consumed food items

	var max1 , max2 , max3 models.MaxCounter

	for food,count := range food_map {

		switch {

		case  max1.Count <= count:
			max3 = max2
			max2 = max1
			max1.Count = count
			max1.Food = food

		case  max2.Count <= count:
			max3 = max2
			max2.Count = count
			max2.Food =  food

		case  max3.Count <= count:
			max3.Count = count
			max3.Food = food

		}

	}

	return max1,max2,max3
}
