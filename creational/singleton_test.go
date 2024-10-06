package creational

import "testing"

func TestGetInstance(testing *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		testing.Error("expected instances to be same")
	}
}
