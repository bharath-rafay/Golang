package main

import (
	"Project/logger"
	"Project/sport"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	file := "/home/bharat/Downloads/in.txt"

	yamlfile := "test.yaml"
	jsonfile := "test.json"

	StructSlice := make([]sport.Details, 0, 3)

	Info, _ := sport.Getinfo(&StructSlice, file)

	//fmt.Println(Info)

	viper.BindEnv("FORMAT")

	env := viper.GetString("FORMAT")

	logger.Info("Environment Variable", zap.String("FORMAT", env))

	if env == "yaml" {
		var YamlCreate sport.MetaYaml

		YamlCreate.Yaml = Info

		sport.CreateFile(&YamlCreate, yamlfile)

	}
	if env == "json" {
		var JsonCreate sport.MetaJson

		JsonCreate.Json = Info

		sport.CreateFile(&JsonCreate, jsonfile)

	}

}
