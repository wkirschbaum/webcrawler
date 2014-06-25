package parser

import (
	"reflect"
	"testing"
)

func TestIdentifyDoubleEscapedHtmlWithout(t *testing.T) {
	parser := Parser{}

	s := `nothing here!`
	result := parser.HasDoubleEscapes(s)
	if result {
		t.Errorf("result should be false")
	}
}

func TestIdentifyDoubleEscapedHtmlWithlt(t *testing.T) {
	parser := Parser{}

	s := `somthing here &lt;a and here`
	result := parser.HasDoubleEscapes(s)
	if !result {
		t.Errorf("result should be true")
	}
}

func TestIdentifyDoubleEscapedHtmlWithAmp(t *testing.T) {
	parser := Parser{}

	s := `somthing here &amp;rarr and here`
	result := parser.HasDoubleEscapes(s)
	if !result {
		t.Errorf("result should be true")
	}
}

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
