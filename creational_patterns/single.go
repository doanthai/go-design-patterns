package creational_patterns

import (
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	Count int
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &Single{}
	}
	return singleInstance
}