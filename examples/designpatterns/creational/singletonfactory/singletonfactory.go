package singletonfactory

import "sync"

type SingletonFactory struct {
	price float64
}

var (
	instance *SingletonFactory
	once     sync.Once
)

func GetInstance() *SingletonFactory {
	once.Do(func() {
		instance = &SingletonFactory{}
	})
	return instance
}

func (sf *SingletonFactory) SetPrice(p float64) {
	sf.price = p
}

func (sf *SingletonFactory) GetPrice() float64 {
	return sf.price
}
