package main

import(
	"fmt"
	"JianZhiOffer_go/2_Singleton/singleton"
)

func main(){
	instance := singleton.GetInstance()
	fmt.Println(instance)
}