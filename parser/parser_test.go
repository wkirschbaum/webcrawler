package parser

import (
	"reflect"
	"testing"
)

func TestParseLinksForLocalLinks(t *testing.T) {
	parser := Parser{BaseUrl: "http://test.com"}

	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	result := parser.ParseLinks(s)

	expected := []string{"http://test.com/foo", "http://test.com/bar/baz"}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestParseLinksIgnoreExternalLinks(t *testing.T) {
	parser := Parser{BaseUrl: "http://test.com"}
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="http://google.com/foo">Foo</a><li></ul>`
	result := parser.ParseLinks(s)

	expected := []string{"http://test.com/foo"}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
