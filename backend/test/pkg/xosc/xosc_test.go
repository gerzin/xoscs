package xosc_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gerzin/xoscs/backend/pkg/xosc"
)

const scenario string = "testdata/dummy_with_parameters.xosc"

func readTestFile(t *testing.T, filename string) []byte {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	filePath := filepath.Join(wd, filename)

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Couldn't find %s: %v", filePath, err)
	}

	return data
}

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

func TestStringerConstraintRule(t *testing.T) {
	var s = xosc.LessOrEqual

	if s.String() != "LessOrEqual" {
		t.Fatalf("Expected lessOrEqual got %s", s)
	}

}

func TestParameterExtraction(t *testing.T) {
	_ = readTestFile(t, scenario)
}
