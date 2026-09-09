package core

import "testing"

func TestParseMarkdownBold(t *testing.T) {
	plain, styles := ParseMarkdown("hello **world**")
	if plain != "hello world" {
		t.Fatalf("plain=%q", plain)
	}
	found := false
	for _, s := range styles {
		if s.Type == "bold" {
			found = true
		}
	}
	if !found {
		t.Fatal("bold style not found")
	}
}

func TestParseHTML(t *testing.T) {
	plain, styles := ParseHTML("<b>hi</b>")
	if plain != "hi" {
		t.Fatalf("plain=%q", plain)
	}
	if len(styles) != 1 || styles[0].Type != "bold" {
		t.Fatalf("styles=%v", styles)
	}
}

func TestParseTextSize(t *testing.T) {
	if ParseTextSize("12") != "12" {
		t.Fatal("size 12 failed")
	}
	if ParseTextSize("abc") != "18" {
		t.Fatal("invalid size should default 18")
	}
}
