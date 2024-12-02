#!/usr/bin/bash

set -e
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cd $SCRIPT_DIR

xgen -i ../pkg/xosc/OpenSCENARIO_1.3.1.xsd -o openscenario_schema -p xosc -l Go
mv openscenario_schema.go ../pkg/xosc/openscenario_schema.go