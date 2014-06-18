package parser

import (
	"reflect"
	"testing"
)

func TestGetLinksFrom(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	result := GetLinksFrom(s)

	expected := []string{"foo", "/bar/baz"}
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}
