package main

import (
	"cube_type_gen/config"
	"cube_type_gen/gen"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func main() {
	generateConfig := flag.Bool("cfg", false, "Generate a config. This is generated at first time. Make sure to adjust configuration!")

	flag.Parse()

	cfgExists := config.Validate()
	if !cfgExists {
		spinner.New().Title("Generating configuration file.").Run()
		config.GenerateDefaultConfig()

		color.HiGreen("Config has been generated, make your modifications and re-run the generator.")
		os.Exit(0)
	}

	if cfgExists && *generateConfig {
		config.GenerateDefaultConfig()
		color.Magenta("Configuration file has been reset to default, make sure to apply your necessary settings.")
		os.Exit(1)
	}

	conf, err := config.Read()
	if err != nil {
		color.HiRed("Could not read or find configuration file. ", err)
		os.Exit(0)
	}

	// make sure the config is not just default barebones. need to make sure its done
	if conf.Prefixes[0].Name == "Placeholder" {
		color.Yellow("Adjust your configuration file to proceed. Re-run this tool once configuration is complete.")
		os.Exit(0)
	}

	if len(conf.Prefixes) == 0 {
		color.HiRed("No prefix list found, are you missing the configuration? Run the generator with -cfg to generate a new configuration file")
		os.Exit(1)
	}

	var generator gen.Generator
	generator.FetchMetadata(conf.CubeUrl)

	spinner.New().Title(fmt.Sprintf("Mapping cubes to prefixes... \n")).Run()

	for i := 0; i < generator.CubeCount; i++ {
		currentCubeName := generator.Metadata.Cubes[i].Name

		for _, value := range conf.Prefixes {
			if value.Name == currentCubeName {

				generator.Metadata.Cubes[i].Name = value.Prefix
				color.Cyan("Gave cube %v the prefix %v \n", currentCubeName, value.Prefix)
			}
		}
	}

	generator.Generate(conf.Output, conf.FileName)

	// Kill the app when complete.
	os.Exit(0)
}
