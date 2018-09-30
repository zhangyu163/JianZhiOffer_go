package main

import "fmt"

func test(testName string, arrary []int, duplicateNum *int, expect bool, errCode error){
	fmt.Println("==============================================")
	fmt.Println(testName)
	res, err := duplicate(arrary, duplicateNum)
	if expect == res && errCode == err{
		fmt.Println("pass")
		if res{
			fmt.Println("the duplicate num is: ", *duplicateNum)
		}
	}
	fmt.Println("==============================================")
}