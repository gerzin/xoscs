package xosc

import (
	"encoding/xml"
	"fmt"
	"os"
)

type ParameterDeclarationI struct {
	Name  string
	Value string
}

func (p ParameterDeclarationI) String() string {
	return fmt.Sprintf("Name: %s, Value: %s", p.Name, p.Value)
}

func ExtractParameters(filePath string) ([]ParameterDeclarationI, error) {
	var parameters []ParameterDeclarationI

	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var scenario OpenSCENARIO
	err = xml.Unmarshal(byteValue, &scenario)
	if err != nil {
		return nil, err
	}

	parDecl := scenario.OpenScenarioCategory.ScenarioDefinition.ParameterDeclarations

	for _, par := range parDecl.ParameterDeclaration {
		parameters = append(parameters, ParameterDeclarationI{
			Name:  par.NameAttr.Parameter,
			Value: par.ValueAttr.Parameter,
		})

	}

	return parameters, nil
}
