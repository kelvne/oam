package main

import (
	"flag"
	"log"

	"github.com/kelvne/oam/v2/mapper"
	"github.com/kelvne/oam/v2/rnp"
)

var (
	modelsFolder string
	baseFile     string
	outputFile   string
)

func init() {
	flag.StringVar(&modelsFolder, "models", ".", "Models directory path")
	flag.StringVar(&baseFile, "base", "", "Path of the base OpenAPI yml file")
	flag.StringVar(&outputFile, "output", "", "Path for the OpenAPI yml file output")
	flag.Parse()
}

func main() {
	parser := rnp.NewParser(modelsFolder)

	if err := parser.ParseFolder(); err != nil {
		log.Fatal(err.Error())
	}

	if err := mapper.MapResourcesToDefinitionFile(parser.Resources, baseFile, outputFile); err != nil {
		log.Fatal(err.Error())
	}
}
