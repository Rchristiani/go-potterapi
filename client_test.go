package potterapi

import (
	"testing"
)

func TestClient(t *testing.T) {
	c := ClientInitialize("123")

	if c.apiKey == "123" {
		t.Log("apiKey is set")
	} else {
		t.Error("apiKey not set")
	}
}
