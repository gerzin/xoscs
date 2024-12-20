package xosc_test

import (
	"os"
	"testing"

	"github.com/gerzin/xoscs/backend/sampler/pkg/xosc"
)

func TestExtractParameters(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		wantErr   bool
		wantCount int
	}{
		{
			name:      "Valid file",
			filePath:  "testdata/valid_scenario.xosc",
			wantErr:   false,
			wantCount: 2,
		},
		{
			name:     "Non-existent file",
			filePath: "testdata/non_existent.xosc",
			wantErr:  true,
		},
		{
			name:     "Invalid XML file",
			filePath: "testdata/invalid_scenario.xosc",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameters, err := xosc.ExtractParameters(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractParameters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(parameters) != tt.wantCount {
				t.Errorf("ExtractParameters() got %v parameters, want %v", len(parameters), tt.wantCount)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.MkdirAll("testdata", os.ModePerm)
	os.WriteFile("testdata/valid_scenario.xosc", []byte(`
		<OpenSCENARIO>
			<ParameterDeclarations>
				<ParameterDeclaration name="param1" type="double" value="0.0"/>
				<ParameterDeclaration name="param2" type="string" value="Ciao!"/>
			</ParameterDeclarations>
		</OpenSCENARIO>
	`), os.ModePerm)
	os.WriteFile("testdata/invalid_scenario.xosc", []byte(`<OpenSCENARIO><ParameterDeclarations>`), os.ModePerm)

	// Run tests
	code := m.Run()

	// Cleanup test data
	os.RemoveAll("testdata")

	os.Exit(code)
}
