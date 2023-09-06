package websocket

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	pool := NewPool()

	if pool == nil {
		t.Errorf("Expected a non-nil Pool, but got nil")
	}

	// Ensure the channels are initialized
	if pool.Register == nil || pool.Unregister == nil || pool.Broadcast == nil {
		t.Errorf("Expected channels to be initialized, but got nil")
	}

	// Ensure the Clients map is initialized
	if pool.Clients == nil {
		t.Errorf("Expected Clients map to be initialized, but got nil")
	}
}

func TestPoolStart(t *testing.T) {
	pool := NewPool()
	go pool.Start()

	if pool.Register == nil || pool.Unregister == nil || pool.Broadcast == nil {
		t.Errorf("Expected channels to be initialized, but got nil")
	}

	if pool.Clients == nil {
		t.Errorf("Expected Clients map to be initialized, but got nil")
	}
}
