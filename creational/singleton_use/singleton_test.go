package use_singleton

import (
	"github.com/adarqui/go_design_pattern/creational/singleton"
	"testing"
)

func TestSingleton(t *testing.T) {

	// cannot refer to unexported name singleton.singleton
	// instance := singleton.singleton{}

	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 != instance2 {
		t.Errorf("instance1 != instance2, not a singleton.\n")
	}

}
