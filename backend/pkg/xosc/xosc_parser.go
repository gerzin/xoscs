package xosc

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"os"
// )

// func ExtractParameters(filename string) ([]OpenScenarioParameter, error) {

// 	file, err := os.Open(filename)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var root OpenSCENARIO

// 	decoder := xml.NewDecoder(file)

// 	if err := decoder.Decode(&root); err != nil {
// 		return nil, fmt.Errorf("failed to decode XML: %w", err)
// 	}

// 	var parameters = root.OpenScenarioCategory.ScenarioDefinition.ParameterDeclarations.ParameterDeclaration

// 	for i, param := range parameters {
// 		fmt.Printf("%d %s\n", i, param.NameAttr.String)
// 	}

// 	return nil, nil

// }
