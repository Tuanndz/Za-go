package worker

import "testing"

func TestNewMessage(t *testing.T) {
	m := NewMessage("hi")
	if m.Text != "hi" {
		t.Fatalf("text=%q", m.Text)
	}
}

func TestMessageStyle(t *testing.T) {
	s := MessageStyle(0, 2, "bold", "", "", false)
	m, ok := s.(map[string]any)
	if !ok {
		t.Fatal("style should be map")
	}
	if m["st"] != "b" {
		t.Fatalf("st=%v", m["st"])
	}
}

func TestMention(t *testing.T) {
	v := Mention("123", 5, 0, false)
	list, ok := v.([]map[string]any)
	if !ok || len(list) != 1 || list[0]["uid"] != "123" {
		t.Fatalf("mention=%v", v)
	}
}
