package main

import "fmt"

func aa() {
	a := []string{"AK5SSA1B38100666",
		"AK5SSA1B38100668",
		"AK5SSA1B38100668",
		"AK5SSA1B38100666",
		"AK5SSA1B38100668",
		"AK5SSA1B38100666",
		"AK5SSA1B38100668",
		"AK5SSA1B38100666"}

	resArr := make([]string, 0)
	tempMap := make(map[string]interface{})
	for _, val := range a {
		_, ok := tempMap[val]
		if !ok {
			resArr = append(resArr, val)
			tempMap[val] = nil
		} else {
			fmt.Println("出现重复")
			break
		}
	}
}
