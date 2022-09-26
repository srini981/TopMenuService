package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileReader() (map[string]int, error){

	//food map to keep track of the food menu items count
	food_map := make(map[string]int)

	// eater_map is used to check whether a food id and menu id is repeated or not
	eater_map:= make(map[string]bool)

	file,err:=os.Open("files/testcase2.log")

	if err !=nil{
		customError:=fmt.Errorf("file does not exists please check the file path")
		return nil,customError
	}

	scanner := bufio.NewScanner(file)

	// reads the file line by line
	for scanner.Scan(){

		if scanner.Text()!="" {

			check := strings.Split(scanner.Text(), " ")
			food_map[check[1]] += 1

			if !eater_map[check[0]+check[1]]  {
				eater_map[check[0]+check[1]] =true
			}else{
				customError:=fmt.Errorf("eater_id and foodmenu_id already exists")
				return nil,customError
 			}
		}
	}
	return food_map,nil
}
