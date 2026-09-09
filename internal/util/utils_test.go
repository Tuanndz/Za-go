package util

import "testing"

func TestNormalizePhone(t *testing.T) {
	cases := map[string]string{
		"0901234567":    "84901234567",
		"+84901234567":  "84901234567",
		"84901234567":   "84901234567",
		" 0901 234 567 ": "84901234567",
	}
	for in, want := range cases {
		if got := NormalizePhone(in); got != want {
			t.Fatalf("NormalizePhone(%q)=%q want %q", in, got, want)
		}
	}
}

func TestAsString(t *testing.T) {
	if AsString(nil) != "" {
		t.Fatal("nil should be empty")
	}
	if AsString(42) != "42" {
		t.Fatal("int convert failed")
	}
	if AsString("  hi ") != "hi" {
		t.Fatal("trim failed")
	}
}

func TestAsInt(t *testing.T) {
	if AsInt("42") != 42 {
		t.Fatal("string int failed")
	}
	if AsInt(3.9) != 3 {
		t.Fatal("float int failed")
	}
	if AsInt(nil) != 0 {
		t.Fatal("nil int should be 0")
	}
}

func TestGetClientMessageType(t *testing.T) {
	if GetClientMessageType("webchat") != 1 {
		t.Fatal("webchat should be 1")
	}
	if GetClientMessageType("chat.photo") != 32 {
		t.Fatal("chat.photo should be 32")
	}
	if GetClientMessageType("unknown-xyz") != 1 {
		t.Fatal("default should be 1")
	}
}

func TestMD5Hex(t *testing.T) {
	if MD5Hex([]byte("hello")) != "5d41402abc4b2a76b9719d911017c592" {
		t.Fatal("md5 mismatch")
	}
}

func TestJSONString(t *testing.T) {
	s := JSONString(map[string]any{"a": 1})
	if s == "" || len(s) < 5 {
		t.Fatal("JSONString empty")
	}
}
