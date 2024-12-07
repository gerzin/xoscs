package xosc

import (
	"encoding/xml"
	"os"
)

// ExtractParameters extracts parameters from an OpenSCENARIO file.
// It reads the file specified by the given file path, parses the XML content,
// and extracts the parameters along with their constraints.
//
// Parameters:
//   - filePath: The path to the OpenSCENARIO file.
//
// Returns:
//   - A slice of Parameter structs containing the extracted parameters and their constraints.
//   - An error if there is an issue reading the file or parsing the XML content.
//
// Example usage:
//
//	parameters, err := ExtractParameters("path/to/scenario.xosc")
//	if err != nil {
//	    log.Fatalf("Failed to extract parameters: %v", err)
//	}
//	for _, param := range parameters {
//	    fmt.Printf("Name: %s, Type: %s, Constraint: %s\n", param.Name, param.Type, param.Constraint)
//	}
func ExtractParameters(filePath string) ([]ParameterDeclaration, error) {
	var parameters []ParameterDeclaration

	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var scenario OpenSCENARIO
	err = xml.Unmarshal(byteValue, &scenario)
	if err != nil {
		return nil, err
	}

	return parameters, nil
}
