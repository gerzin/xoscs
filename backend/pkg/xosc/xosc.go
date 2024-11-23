package xosc

import (
	"os"
)

type OpenScenario struct{}

type OpenScenarioParameterType string

type OpenScenarioParameter struct {
	Name  string
	Type  OpenScenarioParameterType
	Value string
}

func Load(path string) OpenScenario {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

}
