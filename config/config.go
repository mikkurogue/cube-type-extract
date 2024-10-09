package config

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

func GenerateDefaultConfig() {
	config := Configuration{
		CubeUrl:  "http://localhost:4000/cubejs-api",
		Output:   "./",
		FileName: "cubejs-types",
		Prefixes: []Prefix{
			{
				Name:   "Placeholder", // for now we generate the default config with "placeholder" as its first cube, will need to check that this is not named placeholder
				Prefix: "Main",        // This is just something, make sure to explain to users that they should use the -cfg flag before using and edit the cf
			},
		},
		IgnoreErrorCubes: true,
	}

	jsonData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		color.Red("Could not marshal config file:", err)
		os.Exit(0)
	}

	filename := "type-gen-config.json"
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		color.Red("Could not write to file: ", err)
		os.Exit(0)
	}

	color.Green("Successfully created the default configuration file")
}

func Validate() bool {
	_, err := os.Stat("type-gen-config.json")
	return !os.IsNotExist(err)
}

func Read() (*Configuration, error) {

	data, err := os.ReadFile("type-gen-config.json")
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
