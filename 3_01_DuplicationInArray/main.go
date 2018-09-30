package main

func main(){
	var duplicate int
	test("test1 for nill array", []int{}, &duplicate, false, ERR_INVALIDARGUMENT)
	test("test2 for nil duplicate", []int{1, 2, 3, 0}, nil , false, ERR_INVALIDARGUMENT)
	test("test3 for invalid numbers in array", []int{1, 2, 0, 4}, &duplicate, false, ERR_INVALINUMBEROFARRAY)
	test("test4 for invalid numbers in array", []int{1, -2, 0, 0}, &duplicate, false, ERR_INVALINUMBEROFARRAY)
	test("test5 for invalid numbers in array", []int{2, 2, 2, 3, 3, 2, 3, 9, -2}, &duplicate, false, ERR_INVALINUMBEROFARRAY)
	test("test6 2 3", []int{0, 1, 2, 3, 2, 3}, &duplicate, true, nil)
	test("test7 3 1", []int{4, 3, 2, 1, 3, 1}, &duplicate, true, nil)
	test("test8 2", []int{4, 5, 2, 3, 2, 0}, &duplicate, true, nil)
}
