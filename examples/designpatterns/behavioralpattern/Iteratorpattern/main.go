package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Collection struct {
	items []interface{}
}

func (c *Collection) CreateIterator() Iterator {
	return &ConcreteIterator{
		collection: c,
		index:      0,
	}
}

type ConcreteIterator struct {
	collection *Collection
	index      int
}

func (ci *ConcreteIterator) HasNext() bool {
	return ci.index < len(ci.collection.items)
}

func (ci *ConcreteIterator) Next() interface{} {
	item := ci.collection.items[ci.index]
	ci.index++
	return item
}
func main() {
	collection := &Collection{
		items: []interface{}{"A", "B", "C", "D"},
	}

	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}
