package singleton

import(
	"sync"
)

type singleton struct{

}

var Instance *singleton
var once sync.Once

func GetInstance() *singleton{
	once.Do(func(){
		Instance = &singleton{}
	})
	return Instance
}

