package singletonfactory_test

import (
	"singletonfactory"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := singletonfactory.GetInstance()
	instance2 := singletonfactory.GetInstance()

	if instance1 != instance2 {
		t.Error("Expected same instance, got different instances")
	}

	instance1.SetPrice(10)
	if instance2.GetPrice() != 10 {
		t.Error("Expected price to be 10, got", instance2.GetPrice())
	}
}
