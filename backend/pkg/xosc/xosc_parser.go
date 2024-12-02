package xosc

import (
	"os"

	"github.com/gerzin/xoscs/backend/pkg/xosc/schemas/openscenario"
)

func ExtractParameters(filename string) ([]OpenScenarioParameter, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var root openscenario.OpenSCENARIO

}
