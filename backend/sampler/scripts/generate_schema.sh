#!/bin/bash

# Get the directory of the script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Define paths
XSD_FILE="$SCRIPT_DIR/assets/OpenSCENARIO_1.3.1.xsd"
OUTPUT_DIR="$SCRIPT_DIR/../pkg/xosc"

if ! command -v xgen &> /dev/null
then
    echo "xgen could not be found, please install it first."
    exit 1
fi

mkdir -p "$OUTPUT_DIR"

xgen -i "$XSD_FILE" -o "$OUTPUT_DIR/openscenario" -l Go -p xosc

echo "Schema generation complete. Files are located in $OUTPUT_DIR"