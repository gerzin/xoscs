package xosc_test

import (
	"testing"

	"github.com/gerzin/xoscs/backend/pkg/xosc"
)

func TestStringerParameterType(t *testing.T) {
	var s = xosc.String
	var d = xosc.DateTime

	if s.String() != "String" {
		t.Fatalf("Expected String got %s", s)
	}
	if d.String() != "DateTime" {
		t.Fatalf(`Expected DateTime got %s`, s)
	}

}
