package main

import (
	"Project/logger"
	"Project/sport"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if len(os.Args) < 2 {

		logger.Error("Enter the text file as an argument")
		os.Exit(1)

	}

	file := os.Args[1]

	var yamlfile string
	var jsonfile string

	StructSlice := make([]sport.Details, 0, 3)

	Info, err := sport.Getinfo(&StructSlice, file)

	if err != nil {
		logger.Error("Error in returning Slice of struct", zap.Error(err))
	}

	//fmt.Println(Info)

	viper.BindEnv("FORMAT")

	env := viper.GetString("FORMAT")

	logger.Info("Environment Variable", zap.String("FORMAT", env))

	if env == "yaml" {
		fmt.Println("Enter the name of yaml to be created")
		fmt.Scanln(&yamlfile)
		var YamlCreate sport.MetaYaml

		YamlCreate.Yaml = Info

		sport.CreateFile(&YamlCreate, yamlfile)

	}
	if env == "json" {
		fmt.Println("Enter the name of json to be created")
		fmt.Scanln(&jsonfile)

		var JsonCreate sport.MetaJson

		JsonCreate.Json = Info

		sport.CreateFile(&JsonCreate, jsonfile)

	}

}
