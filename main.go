package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	CUBEJS_API_URL = "http://localhost:4000/cubejs-api"
)

type CubeMetadata struct {
	Cubes []Cube `json:"cubes"`
}

type Cube struct {
	Name       string      `json:"name"`
	Dimensions []Dimension `json:"dimensions"`
	Measures   []Measure   `json:"measures"`
}

type Dimension struct {
	Name string `json:"name"`
}

type Measure struct {
	Name string `json:"name"`
}

func main() {
	// Fetch Cube.js metadata
	resp, err := fetchCubejsMetadata()
	if err != nil {
		fmt.Println("Error fetching Cube.js metadata:", err)
		return
	}

	// Parse the metadata JSON
	var metadata CubeMetadata
	if err := json.Unmarshal(resp, &metadata); err != nil {
		fmt.Println("Error parsing Cube.js metadata:", err)
		return
	}

	// Prepare output for TypeScript file
	var output strings.Builder

	for _, cube := range metadata.Cubes {
		var dimensions []string
		var measures []string

		// Extract dimensions and measures without the Cube. prefix
		for _, dimension := range cube.Dimensions {
			dimensionName := extractName(dimension.Name)
			dimensions = append(dimensions, fmt.Sprintf("'%s'", dimensionName))
		}
		for _, measure := range cube.Measures {
			measureName := extractName(measure.Name)
			measures = append(measures, fmt.Sprintf("'%s'", measureName))
		}

		// Generate TypeScript union types for this cube
		cubeName := capitalize(cube.Name)
		dimensionsType := fmt.Sprintf("export type %sDimensions = %s;", cubeName, joinUnion(dimensions))
		measuresType := fmt.Sprintf("export type %sMeasures = %s;", cubeName, joinUnion(measures))

		// Append to the output
		output.WriteString(fmt.Sprintf("%s\n\n%s\n\n", dimensionsType, measuresType))
	}

	// Write to a TypeScript file
	if err := os.WriteFile("cubejs-types.ts", []byte(output.String()), 0644); err != nil {
		fmt.Println("Error writing TypeScript file:", err)
		return
	}

	fmt.Println("Generated cubejs-types.ts with Dimension and Measure union types for each Cube.")
}

func fetchCubejsMetadata() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/meta", CUBEJS_API_URL), nil)
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
