package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

func (g *Generator) FetchMetadata(cubeUrl string) {
	// Fetch Cube.js metadata
	resp, err := fetchCubejsMetadata(cubeUrl)
	if err != nil {
		color.Red("Error fetching Cube.js metadata:", err)
		os.Exit(0)
	}

	// Parse the metadata JSON - Assign metadata to the pointer
	if err := json.Unmarshal(resp, &g.Metadata); err != nil {
		color.Red("Error parsing Cube.js metadata:", err)
		os.Exit(1)
	}

	// Also set the cube count here from metadata.
	g.CubeCount = len(g.Metadata.Cubes)
}

func (g *Generator) Generate(outputDir, filename string, skipErrors bool) {

	// Prepare output for TypeScript file
	var output strings.Builder
	var allDimensionTypes []string
	var allMeasureTypes []string

	fileHeader := fmt.Sprintf("// !! This file is generated by the Cube Type Extracator - Do not modify !!")
	output.WriteString(fmt.Sprintf("%s\n\n", fileHeader))

	for _, cube := range g.Metadata.Cubes {

		if skipErrors == true {
			skip(cube.Name, "error")
		}

		var dimensions []string
		var measures []string

		// Extract dimensions and measures without the Cube. prefix
		for _, dimension := range cube.Dimensions {
			if dimension.Meta.Extractable == true {
				dimensionName := extractName(dimension.Name)
				dimensions = append(dimensions, fmt.Sprintf("'%s'", dimensionName))
			}
		}

		for _, measure := range cube.Measures {
			if measure.Meta.Extractable == true {
				measureName := extractName(measure.Name)
				measures = append(measures, fmt.Sprintf("'%s'", measureName))
			}
		}

		if len(dimensions) != 0 {
			cubeName := capitalize(cube.Name)
			dimensionsTypeName := fmt.Sprintf("%sDimensions", cubeName)
			dimensionsType := fmt.Sprintf("export type %s = %s;", dimensionsTypeName, joinUnion(dimensions))

			output.WriteString(fmt.Sprintf("%s\n\n%s\n\n", fileHeader, dimensionsType))

			allDimensionTypes = append(allDimensionTypes, dimensionsTypeName)
		} else {
			if !containsIgnoreCase(cube.Name, "error") {
				color.Red("Could not generate union type for cube: %v, missing extractable property for dimensions in the schema.", cube.Name)
			}
		}

		if len(measures) != 0 {
			cubeName := capitalize(cube.Name)
			measuresTypeName := fmt.Sprintf("%sMeasures", cubeName)
			measuresType := fmt.Sprintf("export type %s = %s;", measuresTypeName, joinUnion(measures))
			output.WriteString(fmt.Sprintf("%s\n\n", measuresType))

			allMeasureTypes = append(allMeasureTypes, measuresTypeName)

		} else {
			if !containsIgnoreCase(cube.Name, "error") {
				color.Red("Could not generate union type for cube: %v, missing extractable property for measures in the schema", cube.Name)
			}
		}
	}

	// generate all dims and measures as 1 massive union
	if len(allDimensionTypes) > 0 {
		allDimensionsType := fmt.Sprintf("export type AllDimensions = %s;", joinUnion(allDimensionTypes))
		output.WriteString(fmt.Sprintf("%s\n\n", allDimensionsType))
	}

	if len(allMeasureTypes) > 0 {
		allMeasuresType := fmt.Sprintf("export type AllMeasures = %s;", joinUnion(allMeasureTypes))
		output.WriteString(fmt.Sprintf("%s\n\n", allMeasuresType))
	}

	//check if the output dir exists
	if err := os.WriteFile(outputDir+filename+".ts", []byte(output.String()), 0644); err != nil {
		color.Red("Error writing TypeScript file:", err)
		return
	}

	color.Blue("Generated %v.ts with dimension and measure union types.", filename)
}

func fetchCubejsMetadata(cubeUrl string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/meta", cubeUrl), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func extractName(fullName string) string {
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1] // Get the last part after splitting by "."
}

func joinUnion(items []string) string {
	return strings.Join(items, " | ")
}

func capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), substr)
}

func skip(name, checkCase string) {
	if containsIgnoreCase(name, checkCase) {
		color.HiYellow("Skipping cube case '%v' with name '%v'", name, checkCase)
	}
}
