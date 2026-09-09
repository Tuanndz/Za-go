package zago

import "testing"

func TestZaloConstructorNoNetwork(t *testing.T) {
	z, err := Zalo("", "", "test-imei-001", nil, "Mozilla/5.0", false, 0)
	if err != nil {
		t.Fatalf("Zalo() err: %v", err)
	}
	if z == nil {
		t.Fatal("Zalo() nil")
	}
	if z.IsListening() {
		t.Fatal("should not be listening")
	}
}
