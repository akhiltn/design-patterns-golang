package creational

import (
	"sync"
)

type singleton struct {
	count int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
		instance.count++
	})
	return instance
}
