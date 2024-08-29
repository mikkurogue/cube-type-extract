package main

import (
	"bufio"
	"cube_type_gen/gen"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Generator struct {
	CubeCount int
	CubeNames []string
}

var cubeNames string

func main() {

	rename := flag.Bool("rename", false, "Rename the file type prefixes")
	fileName := flag.String("filename", "cubejs-types", "Rename the file without a file extension")
	outputDir := flag.String("o", "./", "Provide an output directory, default is current directory")

	flag.Parse()

	var generator gen.Generator

	generator.FileName = *fileName

	generator.FetchMetadata()

	// check if the flag is set to true, then we start a form to rename the cube metadata
	if *rename == true {
		// map over the generator cube count and assign custom prefixes
		for i := 0; i < generator.CubeCount; i++ {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Enter a prefix for the cube '%v': ", generator.Metadata.Cubes[i].Name)
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")

			generator.Metadata.Cubes[i].Name = text
		}
	}

	generator.IterateToGenerate(*outputDir)

	// Kill the app when complete.
	os.Exit(0)
}
