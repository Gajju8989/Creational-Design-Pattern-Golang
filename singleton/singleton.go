package singleton

import (
	"fmt"
	"sync"
)

var (
	instanse *Singleton
	only     sync.Once
)

type Singleton struct {
}

func GetInstance() *Singleton {

	only.Do(func() {
		instanse = &Singleton{}
		fmt.Print("Singleton Instance Created")
	})

	return instanse
}
