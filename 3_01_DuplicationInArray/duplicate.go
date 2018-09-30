package main

import (
	"errors"
)

var(
	ERR_INVALIDARGUMENT = errors.New("the array or the duplicateNum is nil")
	ERR_INVALINUMBEROFARRAY = errors.New("the numbers of array should be 0 ~ n-1")
)
func duplicate(array []int, duplicateNum *int) (bool, error){
	len := len(array)
	if 0 == len || nil == duplicateNum{
		return false, ERR_INVALIDARGUMENT
	}
	for _, i := range array{
		if i < 0 || i > len - 1{
			*duplicateNum = -1
			return false, ERR_INVALINUMBEROFARRAY
		}
	}
	for i, _ := range array{
		for{
			if array[i] != i {
				if array[i] == array[array[i]]{
					*duplicateNum = array[i]
					return true, nil
				}
				array[i], array[array[i]] = array[array[i]], array[i]
				//temp := array[i]
				//array[i] = array[array[i]]
				//array[temp] = temp
			}else{
				break
			}
		}
	}
	*duplicateNum = -1
	return false, nil
}